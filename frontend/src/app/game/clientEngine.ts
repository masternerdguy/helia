import { WsService } from './ws.service';
import { ServerJoinBody } from './wsModels/bodies/join';
import { GameMessage, MessageTypes } from './wsModels/gameMessage';
import { Player, TargetType } from './engineModels/player';
import { System } from './engineModels/system';
import { Ship } from './engineModels/ship';
import { Camera } from './engineModels/camera';
import { ServerGlobalUpdateBody } from './wsModels/bodies/globalUpdate';
import { Backplate } from './procedural/backplate/backplate';
import { ClientNavClick } from './wsModels/bodies/navClick';
import { angleBetween, magnitude } from './engineMath';
import { Star } from './engineModels/star';
import { Planet } from './engineModels/planet';
import { GDIWindow } from './gdi/base/gdiWindow';
import { Station } from './engineModels/station';
import { GDIStyle } from './gdi/base/gdiStyle';
import { Jumphole } from './engineModels/jumphole';
import { ShipStatusWindow } from './gdi/windows/shipStatusWindow';
import { ServerCurrentShipUpdate } from './wsModels/bodies/currentShipUpdate';
import { TargetInteractionWindow } from './gdi/windows/targetInterationWindow';
import { OverviewWindow } from './gdi/windows/overviewWindow';
import { ModuleEffect } from './engineModels/moduleEffect';
import { PointEffect } from './engineModels/pointEffect';
import { WindowManager } from './gdi/windows/windowManager';
import { ShipFittingWindow } from './gdi/windows/shipFittingWindow';
import { ServerContainerView } from './wsModels/bodies/containerView';
import { Container } from './engineModels/container';
import { ServerErrorMessage } from './wsModels/bodies/errorMessage';
import { Asteroid } from './engineModels/asteroid';
import { OrdersMarketWindow } from './gdi/windows/ordersMarketWindow';
import { ServerOpenSellOrdersUpdate } from './wsModels/bodies/openSellOrdersUpdate';
import { PushErrorWindow } from './gdi/windows/pushErrorWindow';
import { IndustrialMarketWindow } from './gdi/windows/industrialMarketWindow';
import { ServerIndustrialOrdersUpdate } from './wsModels/bodies/industrialOrdersUpdate';
import { ServerFactionUpdate } from './wsModels/bodies/factionUpdate';
import {
  UpdateFactionCache,
  UpdatePlayerFactionRelationshipCache,
} from './wsModels/shared';
import { StarMapWindow } from './gdi/windows/starMapWindow';
import {
  ServerStarMapUpdate,
  UnwrappedStarMapData,
} from './wsModels/bodies/viewStarMap';
import { ReputationSheetWindow } from './gdi/windows/reputationSheetWindow';
import { ServerPlayerFactionUpdate } from './wsModels/bodies/playerFactionUpdate';
import { PropertySheetWindow } from './gdi/windows/propertySheetWindow';
import { ServerPropertyUpdate } from './wsModels/bodies/propertyUpdate';
import { Missile } from './engineModels/missile';
import { SystemChatWindow } from './gdi/windows/systemChatWindow';
import { ServerInfoMessage } from './wsModels/bodies/infoMessage';
import { PushInfoWindow } from './gdi/windows/pushInfoWindow';
import { ExperienceSheetWindow } from './gdi/windows/experienceSheetWindow';
import { ServerExperienceUpdate } from './wsModels/bodies/experienceUpdate';
import { ClientGlobalAck } from './wsModels/bodies/globalAck';
import { SchematicRunsWindow } from './gdi/windows/schematicRunsWindow';
import { ServerSchematicRunsUpdate } from './wsModels/bodies/schematicRunsUpdate';
import { Wreck } from './engineModels/wreck';
import { ServerApplicationsUpdate } from './wsModels/bodies/applicationsUpdate';
import { ServerMembersUpdate } from './wsModels/bodies/membersUpdate';
import { ActionReportsWindow } from './gdi/windows/actionReportsWindow';
import { ServerActionReportsPage } from './wsModels/bodies/viewActionReportsPage';
import * as ClipboardJS from 'clipboard';
import { ServerDockedUsersUpdate } from './wsModels/bodies/dockedUsersUpdate';
import { Outpost } from './engineModels/outpost';
import { Artifact } from './engineModels/artifact';

class EngineSack {
  constructor() {}

  // player
  player: Player;

  // camera
  camera: Camera;

  // foreground graphics
  gfx: HTMLCanvasElement;
  ctx: any;

  // backplate graphics
  backplateCanvas: HTMLCanvasElement;
  backplateRenderer: Backplate;

  // ui elements
  pushErrorWindow: PushErrorWindow;
  pushInfoWindow: PushInfoWindow;
  shipStatusWindow: ShipStatusWindow;
  targetInteractionWindow: TargetInteractionWindow;
  overviewWindow: OverviewWindow;
  shipFittingWindow: ShipFittingWindow;
  ordersMarketWindow: OrdersMarketWindow;
  industrialMarketWindow: IndustrialMarketWindow;
  starMapWindow: StarMapWindow;
  reputationSheetWindow: ReputationSheetWindow;
  propertySheetWindow: PropertySheetWindow;
  experienceSheetWindow: ExperienceSheetWindow;
  schematicRunsWindow: SchematicRunsWindow;
  systemChatWindow: SystemChatWindow;
  actionReportsWindow: ActionReportsWindow;
  lastShiftDown: number;

  tpsTail: number[] = [];

  windows: GDIWindow[];
  windowManager: WindowManager;
  alternateFrame: boolean;

  // client-server communication
  wsSvc: WsService;

  // tpf
  lastFrameTime: number;
  lastSyncTime: number;
  tpf: number;
  tps: number;
  tpi: number;

  reloading = false;

  // mouse cache
  mouseX = 0;
  mouseY = 0;
}

const engineSack: EngineSack = new EngineSack();

export function clientStart(
  wsService: WsService,
  gameCanvas: HTMLCanvasElement,
  backCanvas: HTMLCanvasElement,
  sid: string,
) {
  // set canvases to initial width
  const clientWidth = document.documentElement.clientWidth;
  const clientHeight = document.documentElement.clientHeight;

  gameCanvas.width = clientWidth;
  gameCanvas.height = clientHeight;

  backCanvas.width = clientWidth;
  backCanvas.height = clientHeight;

  // initialize player
  engineSack.player = new Player();

  // initialize camera
  engineSack.camera = new Camera(gameCanvas.width, gameCanvas.height, 1);

  // initialize backplate
  engineSack.backplateRenderer = new Backplate(backCanvas);

  // initialize window manager
  engineSack.windowManager = new WindowManager();
  engineSack.windowManager.preinit(gameCanvas.height, (w: GDIWindow) => {
    // move latest toggled window to top
    engineSack.windows = [
      w,
      ...engineSack.windows.filter((item) => item !== w),
    ];
  });

  engineSack.windowManager.setX(0);
  engineSack.windowManager.setY(Number.NEGATIVE_INFINITY);
  engineSack.windowManager.initialize();

  // initialize ui windows
  engineSack.pushErrorWindow = new PushErrorWindow();
  engineSack.pushErrorWindow.initialize();
  engineSack.pushErrorWindow.pack();
  engineSack.pushErrorWindow.setX(
    gameCanvas.width / 2 - engineSack.pushErrorWindow.getWidth() / 2,
  );
  engineSack.pushErrorWindow.setY(
    gameCanvas.height / 2 - engineSack.pushErrorWindow.getHeight() / 2,
  );

  engineSack.pushInfoWindow = new PushInfoWindow();
  engineSack.pushInfoWindow.initialize();
  engineSack.pushInfoWindow.pack();
  engineSack.pushInfoWindow.setX(
    gameCanvas.width / 2 - engineSack.pushInfoWindow.getWidth() / 2,
  );
  engineSack.pushInfoWindow.setY(
    gameCanvas.height / 2 - engineSack.pushInfoWindow.getHeight() / 2,
  );

  engineSack.shipStatusWindow = new ShipStatusWindow();
  engineSack.shipStatusWindow.setX(engineSack.windowManager.getWidth());
  engineSack.shipStatusWindow.setY(0);
  engineSack.shipStatusWindow.initialize();
  engineSack.shipStatusWindow.pack();
  engineSack.shipStatusWindow.setWsService(wsService);
  engineSack.shipStatusWindow.setPlayer(engineSack.player);

  engineSack.shipFittingWindow = new ShipFittingWindow();
  engineSack.shipFittingWindow.setX(375);
  engineSack.shipFittingWindow.setY(375);
  engineSack.shipFittingWindow.initialize();
  engineSack.shipFittingWindow.pack();
  engineSack.shipFittingWindow.setWsService(wsService);
  engineSack.shipFittingWindow.setPlayer(engineSack.player);

  engineSack.ordersMarketWindow = new OrdersMarketWindow();
  engineSack.ordersMarketWindow.setX(425);
  engineSack.ordersMarketWindow.setY(425);
  engineSack.ordersMarketWindow.initialize();
  engineSack.ordersMarketWindow.pack();
  engineSack.ordersMarketWindow.setWsService(wsService);
  engineSack.ordersMarketWindow.setPlayer(engineSack.player);

  engineSack.industrialMarketWindow = new IndustrialMarketWindow();
  engineSack.industrialMarketWindow.setX(325);
  engineSack.industrialMarketWindow.setY(325);
  engineSack.industrialMarketWindow.initialize();
  engineSack.industrialMarketWindow.pack();
  engineSack.industrialMarketWindow.setWsService(wsService);
  engineSack.industrialMarketWindow.setPlayer(engineSack.player);

  engineSack.targetInteractionWindow = new TargetInteractionWindow();
  engineSack.targetInteractionWindow.initialize();
  engineSack.targetInteractionWindow.setWsSvc(wsService);
  engineSack.targetInteractionWindow.setX(
    gameCanvas.width - engineSack.targetInteractionWindow.getWidth(),
  );
  engineSack.targetInteractionWindow.setY(0);
  engineSack.targetInteractionWindow.pack();

  engineSack.overviewWindow = new OverviewWindow();
  engineSack.overviewWindow.setHeight(gameCanvas.height / 2);
  engineSack.overviewWindow.initialize();
  engineSack.overviewWindow.setX(
    gameCanvas.width - engineSack.overviewWindow.getWidth(),
  );
  engineSack.overviewWindow.setY(
    engineSack.targetInteractionWindow.getY() +
      engineSack.targetInteractionWindow.getHeight() +
      GDIStyle.windowHandleHeight,
  );
  engineSack.overviewWindow.pack();

  engineSack.starMapWindow = new StarMapWindow();
  engineSack.starMapWindow.setX(455);
  engineSack.starMapWindow.setY(455);
  engineSack.starMapWindow.initialize();
  engineSack.starMapWindow.pack();
  engineSack.starMapWindow.setWsService(wsService);
  engineSack.starMapWindow.setPlayer(engineSack.player);

  engineSack.reputationSheetWindow = new ReputationSheetWindow();
  engineSack.reputationSheetWindow.setX(355);
  engineSack.reputationSheetWindow.setY(355);
  engineSack.reputationSheetWindow.initialize();
  engineSack.reputationSheetWindow.pack();

  engineSack.propertySheetWindow = new PropertySheetWindow();
  engineSack.propertySheetWindow.setX(365);
  engineSack.propertySheetWindow.setY(365);
  engineSack.propertySheetWindow.initialize();
  engineSack.propertySheetWindow.pack();
  engineSack.propertySheetWindow.setPlayer(engineSack.player);

  engineSack.systemChatWindow = new SystemChatWindow();
  engineSack.systemChatWindow.initialize();
  engineSack.systemChatWindow.setX(engineSack.windowManager.getWidth());
  engineSack.systemChatWindow.setY(
    gameCanvas.height -
      engineSack.systemChatWindow.getHeight() -
      GDIStyle.windowHandleHeight -
      1,
  );
  engineSack.systemChatWindow.pack();

  engineSack.experienceSheetWindow = new ExperienceSheetWindow();
  engineSack.experienceSheetWindow.setX(395);
  engineSack.experienceSheetWindow.setY(395);
  engineSack.experienceSheetWindow.initialize();
  engineSack.experienceSheetWindow.setWsService(engineSack.wsSvc);
  engineSack.experienceSheetWindow.pack();

  engineSack.schematicRunsWindow = new SchematicRunsWindow();
  engineSack.schematicRunsWindow.setX(410);
  engineSack.schematicRunsWindow.setY(410);
  engineSack.schematicRunsWindow.initialize();
  engineSack.schematicRunsWindow.setWsService(engineSack.wsSvc);
  engineSack.schematicRunsWindow.pack();

  engineSack.actionReportsWindow = new ActionReportsWindow();
  engineSack.actionReportsWindow.setX(410);
  engineSack.actionReportsWindow.setY(410);
  engineSack.actionReportsWindow.initialize();
  engineSack.actionReportsWindow.setWsService(engineSack.wsSvc);
  engineSack.actionReportsWindow.pack();

  // link windows to window manager
  engineSack.windowManager.manageWindow(engineSack.overviewWindow, '☀');
  engineSack.windowManager.manageWindow(engineSack.shipStatusWindow, '☍');
  engineSack.windowManager.manageWindow(
    engineSack.targetInteractionWindow,
    '☉',
  );
  engineSack.windowManager.manageWindow(engineSack.shipFittingWindow, 'Ʌ');
  engineSack.windowManager.manageWindow(engineSack.ordersMarketWindow, '₪');
  engineSack.windowManager.manageWindow(
    engineSack.industrialMarketWindow,
    '⚙',
  );
  engineSack.windowManager.manageWindow(engineSack.starMapWindow, '⊞');
  engineSack.windowManager.manageWindow(engineSack.reputationSheetWindow, '❉');
  engineSack.windowManager.manageWindow(engineSack.propertySheetWindow, '⬢');
  engineSack.windowManager.manageWindow(engineSack.systemChatWindow, '⋉');
  engineSack.windowManager.manageWindow(engineSack.experienceSheetWindow, '✇');
  engineSack.windowManager.manageWindow(engineSack.schematicRunsWindow, '⨻');
  engineSack.windowManager.manageWindow(engineSack.actionReportsWindow, 'ⓚ');

  // cache windows for simpler updating and rendering
  engineSack.windows = [
    engineSack.pushErrorWindow,
    engineSack.pushInfoWindow,
    engineSack.shipStatusWindow,
    engineSack.targetInteractionWindow,
    engineSack.overviewWindow,
    engineSack.shipFittingWindow,
    engineSack.ordersMarketWindow,
    engineSack.industrialMarketWindow,
    engineSack.starMapWindow,
    engineSack.reputationSheetWindow,
    engineSack.propertySheetWindow,
    engineSack.systemChatWindow,
    engineSack.experienceSheetWindow,
    engineSack.schematicRunsWindow,
    engineSack.actionReportsWindow,
    engineSack.windowManager,
  ];

  // store globals
  engineSack.gfx = gameCanvas;
  engineSack.ctx = gameCanvas.getContext('2d');
  engineSack.backplateCanvas = backCanvas;
  engineSack.wsSvc = wsService;
  engineSack.player.sid = sid;

  // connect
  wsService.connect(sid, (d, ws) => {
    if (d.type === MessageTypes.Join) {
      handleJoin(d);
    } else if (d.type === MessageTypes.Update) {
      handleGlobalUpdate(d);
    } else if (d.type === MessageTypes.CurrentShipUpdate) {
      handleCurrentShipUpdate(d);
    } else if (d.type === MessageTypes.CargoBayUpdate) {
      handleCargoBayUpdate(d);
    } else if (d.type === MessageTypes.PushError) {
      handleErrorMessageFromServer(d);
    } else if (d.type === MessageTypes.PushInfo) {
      handleInfoMessageFromServer(d);
    } else if (d.type === MessageTypes.OpenSellOrdersUpdate) {
      handleOpenSellOrdersUpdateMessageFromServer(d);
    } else if (d.type === MessageTypes.IndustrialOrdersUpdate) {
      handleIndustrialOrdersUpdateMessageFromServer(d);
    } else if (d.type == MessageTypes.FactionUpdate) {
      handleFactionUpdate(d);
    } else if (d.type == MessageTypes.StarMapUpdate) {
      handleStarMapUpdate(d);
    } else if (d.type == MessageTypes.PlayerFactionUpdate) {
      handlePlayerFactionUpdate(d);
    } else if (d.type == MessageTypes.PropertyUpdate) {
      handlePropertyUpdate(d);
    } else if (d.type == MessageTypes.ExperienceUpdate) {
      handleExperienceUpdate(d);
    } else if (d.type == MessageTypes.SchematicRunsUpdate) {
      handleSchematicRunsUpdate(d);
    } else if (d.type == MessageTypes.ApplicationsUpdate) {
      handleApplicationsUpdate(d);
    } else if (d.type == MessageTypes.MembersUpdate) {
      handleMembersUpdate(d);
    } else if (d.type == MessageTypes.ActionReportsPage) {
      handleActionReportsPage(d);
    } else if (d.type == MessageTypes.ActionReportDetail) {
      handleActionReportDetails(d);
    } else if (d.type == MessageTypes.DockedUsersUpdate) {
      handleDockedUsersUpdate(d);
    }
  });
}

function handleJoin(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerJoinBody;

  // stash world welcome
  engineSack.player.uid = msg.currentShipInfo.uid;
  engineSack.player.currentShip = new Ship(msg.currentShipInfo);
  engineSack.player.currentSystem = new System(msg.currentSystemInfo);

  // add click event handler
  engineSack.gfx.addEventListener('click', (event) => {
    // handle event
    handleClick(event.x, event.y);
  });

  // add mouse move event handler
  engineSack.gfx.addEventListener('mousemove', (event) => {
    // handle event
    handleMouseMove(event.x, event.y);
  });

  // add mouse scroll event handler
  engineSack.gfx.addEventListener('wheel', (event) => {
    // get scroll direction
    const delta = Math.sign(event.deltaY);

    // handle event
    handleScroll(delta);
  });

  // add key down event handler
  document.addEventListener('keydown', (event) => {
    // handle event
    handleKeydown(event.key);
  });

  // add window resize listener
  window.addEventListener('resize', () => {
    handleWindowResize();
  });

  // hide initially hidden windows
  engineSack.shipFittingWindow.setHidden(true);
  engineSack.ordersMarketWindow.setHidden(true);
  engineSack.industrialMarketWindow.setHidden(true);
  engineSack.pushErrorWindow.setHidden(true);
  engineSack.pushInfoWindow.setHidden(true);
  engineSack.starMapWindow.setHidden(true);
  engineSack.reputationSheetWindow.setHidden(true);
  engineSack.propertySheetWindow.setHidden(true);
  engineSack.experienceSheetWindow.setHidden(true);
  engineSack.schematicRunsWindow.setHidden(true);
  engineSack.actionReportsWindow.setHidden(true);

  // start game loop
  engineSack.lastFrameTime = Date.now();
  engineSack.lastSyncTime = Date.now();
  engineSack.tpf = 0;

  setTimeout(clientLoop, 0);
  setTimeout(interpolate, 0);
}

function handleGlobalUpdate(d: GameMessage) {
  // store sync time
  const now = Date.now();

  engineSack.tps = now - engineSack.lastSyncTime;
  engineSack.lastSyncTime = now;

  if (engineSack.tpsTail.length >= 5) {
    engineSack.tpsTail = engineSack.tpsTail.reverse();
    engineSack.tpsTail.pop();
    engineSack.tpsTail = engineSack.tpsTail.reverse();
  }

  engineSack.tpsTail.push(engineSack.tps);

  // parse body
  const msg = JSON.parse(d.body) as ServerGlobalUpdateBody;

  // store ack token
  engineSack.wsSvc.ackToken = msg.token;

  setTimeout(() => {
    // send global update ack response
    const ack = new ClientGlobalAck();
    ack.sid = engineSack.wsSvc.sid;
    ack.sysId = engineSack.player.currentSystem.id;
    ack.token = engineSack.wsSvc.ackToken;

    engineSack.wsSvc.sendMessage(MessageTypes.GlobalAck, ack);
  }, 0);

  // system switch or update check
  if (msg.currentSystemInfo.id !== engineSack.player.currentSystem.id) {
    // reinitialize system cache
    engineSack.player.currentSystem = new System(msg.currentSystemInfo);

    // clear target selection
    engineSack.player.currentTargetID = undefined;
    engineSack.player.currentTargetType = undefined;
    engineSack.overviewWindow.selectedItemID = undefined;
    engineSack.overviewWindow.selectedItemType = undefined;
    engineSack.targetInteractionWindow.setTarget(undefined, undefined);

    // clear camera look
    engineSack.targetInteractionWindow.setCameraLook(undefined, undefined);
  } else {
    // fix empty arrays in incoming data
    if (!msg.planets || msg.planets == null) {
      msg.planets = [];
    }

    if (!msg.artifacts || msg.artifacts == null) {
      msg.artifacts = [];
    }

    if (!msg.stations || msg.stations == null) {
      msg.stations = [];
    }

    if (!msg.outposts || msg.outposts == null) {
      msg.outposts = [];
    }

    if (!msg.ships || msg.ships == null) {
      msg.ships = [];
    }

    if (!msg.jumpholes || msg.jumpholes == null) {
      msg.jumpholes = [];
    }

    if (!msg.stars || msg.stars == null) {
      msg.stars = [];
    }

    if (!msg.asteroids || msg.asteroids == null) {
      msg.asteroids = [];
    }

    if (!msg.newModuleEffects || msg.newModuleEffects == null) {
      msg.newModuleEffects = [];
    }

    if (!msg.newPointEffects || msg.newPointEffects == null) {
      msg.newPointEffects = [];
    }

    if (!msg.missiles || msg.missiles == null) {
      msg.missiles = [];
    }

    if (!msg.wrecks || msg.wrecks == null) {
      msg.wrecks = [];
    }

    if (!msg.systemChat || msg.systemChat == null) {
      msg.systemChat = [];
    }

    // update system
    engineSack.player.currentSystem.id = msg.currentSystemInfo.id;
    engineSack.player.currentSystem.systemName =
      msg.currentSystemInfo.systemName;

    // update ships
    for (const sh of msg.ships) {
      let match = false;

      // find ship in memory
      for (const sm of engineSack.player.currentSystem.ships) {
        if (sh.id === sm.id) {
          match = true;

          // todo: update this when players are eventually able to command multiple ships at once
          sm.isPlayer = false;
          sm.isTargeted = false;

          // sync ship in memory
          sm.sync(sh);

          // is this the player ship?
          if (sm.id === engineSack.player.currentShip.id) {
            // flag as player's ship
            sm.isPlayer = true;

            // update player ship cache
            engineSack.player.currentShip.sync(sh);
            engineSack.player.currentShip.isPlayer = true;
          }

          // exit loop
          break;
        }
      }

      if (!match) {
        // add ship to memory
        engineSack.player.currentSystem.ships.push(new Ship(sh));
      }
    }

    // update missiles
    for (const sh of msg.missiles) {
      let match = false;

      // find missile in memory
      for (const sm of engineSack.player.currentSystem.missiles) {
        if (sh.id === sm.id) {
          match = true;

          // sync missile in memory
          sm.sync(sh);

          // exit loop
          break;
        }
      }

      if (!match) {
        // add missile to memory
        engineSack.player.currentSystem.missiles.push(new Missile(sh));
      }
    }

    // update wrecks
    for (const sh of msg.wrecks) {
      let match = false;

      // find wreck in memory
      for (const sm of engineSack.player.currentSystem.wrecks) {
        if (sh.id === sm.id) {
          match = true;

          // sync wreck in memory
          sm.sync(sh);

          // exit loop
          break;
        }
      }

      if (!match) {
        // add wreck to memory
        engineSack.player.currentSystem.wrecks.push(new Wreck(sh));
      }
    }

    // update stars
    for (const st of msg.stars) {
      let match = false;

      // find star in memory
      for (const sm of engineSack.player.currentSystem.stars) {
        if (st.id === sm.id) {
          match = true;

          // sync star in memory
          sm.sync(st);

          // exit loop
          break;
        }
      }

      if (!match) {
        // add star to memory
        engineSack.player.currentSystem.stars.push(new Star(st));
      }
    }

    // update planets
    for (const p of msg.planets) {
      let match = false;

      // find planet in memory
      for (const sm of engineSack.player.currentSystem.planets) {
        if (p.id === sm.id) {
          match = true;

          // sync planet in memory
          sm.sync(p);

          // exit loop
          break;
        }
      }

      if (!match) {
        // add planet to memory
        engineSack.player.currentSystem.planets.push(new Planet(p));
      }
    }

    // update artifacts
    for (const p of msg.artifacts) {
      let match = false;

      // find artifact in memory
      for (const sm of engineSack.player.currentSystem.artifacts) {
        if (p.id === sm.id) {
          match = true;

          // sync artifact in memory
          sm.sync(p);

          // exit loop
          break;
        }
      }

      if (!match) {
        // add artifact to memory
        engineSack.player.currentSystem.artifacts.push(new Artifact(p));
      }
    }

    // update asteroids
    for (const p of msg.asteroids) {
      let match = false;

      // find asteroid in memory
      for (const sm of engineSack.player.currentSystem.asteroids) {
        if (p.id === sm.id) {
          match = true;

          // sync asteroid in memory
          sm.sync(p);

          // exit loop
          break;
        }
      }

      if (!match) {
        // add asteroid to memory
        engineSack.player.currentSystem.asteroids.push(new Asteroid(p));
      }
    }

    // update jumpholes
    for (const j of msg.jumpholes) {
      let match = false;

      // find jumphole in memory
      for (const sm of engineSack.player.currentSystem.jumpholes) {
        if (j.id === sm.id) {
          match = true;

          // sync jumphole in memory
          sm.sync(j);

          // exit loop
          break;
        }
      }

      if (!match) {
        // add jumphole to memory
        engineSack.player.currentSystem.jumpholes.push(new Jumphole(j));
      }
    }

    // update npc stations
    for (const p of msg.stations) {
      let match = false;

      // find station in memory
      for (const sm of engineSack.player.currentSystem.stations) {
        if (p.id === sm.id) {
          match = true;

          // sync station in memory
          sm.sync(p);

          // exit loop
          break;
        }
      }

      if (!match) {
        // add station to memory
        engineSack.player.currentSystem.stations.push(new Station(p));
      }

      // note: npc stations are indestructible for gameplay reasons, but the player owned equivalent "outposts" are not!
    }

    // update outposts
    for (const sh of msg.outposts) {
      let match = false;

      // find outpost in memory
      for (const sm of engineSack.player.currentSystem.outposts) {
        if (sh.id === sm.id) {
          match = true;

          // sync outpost in memory
          sm.sync(sh);

          // exit loop
          break;
        }
      }

      if (!match) {
        // add outpost to memory
        engineSack.player.currentSystem.outposts.push(new Outpost(sh));
      }
    }

    // start new module effects
    for (const ef of msg.newModuleEffects) {
      // copy values
      const effect = new ModuleEffect(ef, engineSack.player);

      // append
      engineSack.player.currentSystem.moduleEffects.push(effect);
    }

    // start new point effects
    for (const ef of msg.newPointEffects) {
      // copy values
      const effect = new PointEffect(ef, engineSack.player);

      // append
      engineSack.player.currentSystem.pointEffects.push(effect);
    }

    // get rid of expired module effects
    const keepModuleEffects: ModuleEffect[] = [];

    for (const ef of engineSack.player.currentSystem.moduleEffects) {
      if (!ef.finished) {
        keepModuleEffects.push(ef);
      }
    }

    engineSack.player.currentSystem.moduleEffects = keepModuleEffects;

    // get rid of expired point effects
    const keepPointEffects: PointEffect[] = [];

    for (const ef of engineSack.player.currentSystem.pointEffects) {
      if (!ef.finished) {
        keepPointEffects.push(ef);
      }
    }

    engineSack.player.currentSystem.pointEffects = keepPointEffects;
  }

  // update overview window
  engineSack.overviewWindow.sync(engineSack.player);

  // update system chat window
  engineSack.systemChatWindow.sync(msg.systemChat);
}

function handleErrorMessageFromServer(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerErrorMessage;

  // move push error window to top
  engineSack.windows = [
    engineSack.pushErrorWindow,
    ...engineSack.windows.filter((item) => item !== engineSack.pushErrorWindow),
  ];

  // show the push error window
  engineSack.pushErrorWindow.setHidden(false);
  engineSack.pushErrorWindow.setText(msg.message);
}

function handleInfoMessageFromServer(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerInfoMessage;

  // move push info window to top
  engineSack.windows = [
    engineSack.pushInfoWindow,
    ...engineSack.windows.filter((item) => item !== engineSack.pushInfoWindow),
  ];

  // show the push info window
  engineSack.pushInfoWindow.setHidden(false);
  engineSack.pushInfoWindow.setText(msg.message);
}

function handleCargoBayUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerContainerView;

  // null check
  if (!msg.items) {
    msg.items = [];
  }

  // update current cargo view cache
  const vw = new Container(msg);
  engineSack.player.currentCargoView = vw;
}

function handleStarMapUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerStarMapUpdate;

  // null check
  if (msg.cachedMapData) {
    // update starmap cache
    engineSack.player.currentStarMap = new UnwrappedStarMapData(
      msg.cachedMapData,
    );
  }
}

function handleOpenSellOrdersUpdateMessageFromServer(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerOpenSellOrdersUpdate;

  // null check
  if (!msg.orders) {
    msg.orders = [];
  }

  // update sell orders window
  engineSack.ordersMarketWindow.syncOpenSellOrders(msg);
}

function handleIndustrialOrdersUpdateMessageFromServer(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerIndustrialOrdersUpdate;

  // null check
  if (!msg.inSilos) {
    msg.inSilos = [];
  }

  if (!msg.outSilos) {
    msg.outSilos = [];
  }

  // update industrial orders window
  engineSack.industrialMarketWindow.syncIndustrialOrders(msg);
}

function handleCurrentShipUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerCurrentShipUpdate;

  // update current ship cache
  if (msg.currentShipInfo.id !== engineSack.player.currentShip.id) {
    // player has changed ships somehow
    engineSack.player.currentShip = new Ship(msg.currentShipInfo);
  }

  engineSack.player.currentShip.sync(msg.currentShipInfo);

  // make sure it's synced on the ship list as well
  for (const sh of engineSack.player.currentSystem.ships) {
    if (sh.id == engineSack.player.currentShip.id) {
      // don't accept new position
      const cx = sh.x;
      const cy = sh.y;

      msg.currentShipInfo.x = cx;
      msg.currentShipInfo.y = cy;

      // sync everything else
      sh.sync(msg.currentShipInfo);
      break;
    }
  }

  // dock check
  if (!!msg.currentShipInfo.dockedAtStationID) {
    // store docked at id as target
    engineSack.player.currentTargetID = msg.currentShipInfo.dockedAtStationID;

    // check if docked at station
    if (
      engineSack.player.currentSystem.stations.filter(
        (x) => x.id == msg.currentShipInfo.dockedAtStationID,
      ).length == 1
    ) {
      engineSack.player.currentTargetType = TargetType.Station;
    }

    // check if docked at outpost
    if (
      engineSack.player.currentSystem.outposts.filter(
        (x) => x.id == msg.currentShipInfo.dockedAtStationID,
      ).length == 1
    ) {
      engineSack.player.currentTargetType = TargetType.Outpost;
    }

    // store target on overview window as well
    engineSack.overviewWindow.selectedItemID =
      engineSack.player.currentTargetID;
    engineSack.overviewWindow.selectedItemType =
      engineSack.player.currentTargetType;
  }

  // update status window
  engineSack.shipStatusWindow.setShip(engineSack.player.currentShip);

  // update fitting window
  engineSack.shipFittingWindow.setPlayer(engineSack.player);
  engineSack.shipFittingWindow.setWsService(engineSack.wsSvc);

  // update orders market window
  engineSack.ordersMarketWindow.setPlayer(engineSack.player);
  engineSack.ordersMarketWindow.setWsService(engineSack.wsSvc);

  // update industrial market window
  engineSack.industrialMarketWindow.setPlayer(engineSack.player);
  engineSack.industrialMarketWindow.setWsService(engineSack.wsSvc);

  // update property window
  engineSack.propertySheetWindow.setWsService(engineSack.wsSvc);
  engineSack.propertySheetWindow.setPlayer(engineSack.player);

  // update system chat window
  engineSack.systemChatWindow.setWsService(engineSack.wsSvc);

  // update experience window
  engineSack.experienceSheetWindow.setWsService(engineSack.wsSvc);

  // update schematic runs window
  engineSack.schematicRunsWindow.setWsService(engineSack.wsSvc);

  // update reputation sheet window
  engineSack.reputationSheetWindow.setWsService(engineSack.wsSvc);
  engineSack.reputationSheetWindow.setPlayer(engineSack.player);

  // update action reports window
  engineSack.actionReportsWindow.setWsService(engineSack.wsSvc);
}

function handleFactionUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerFactionUpdate;

  // update faction cache dictionary
  UpdateFactionCache(msg.factions);
}

function handlePlayerFactionUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerPlayerFactionUpdate;

  // update faction cache dictionary
  UpdatePlayerFactionRelationshipCache(msg.factions);
}

function handlePropertyUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerPropertyUpdate;

  // fix missing lists
  if (!msg.ships) {
    msg.ships = [];
  }

  if (!msg.outposts) {
    msg.outposts = [];
  }

  // update property sheet window
  engineSack.propertySheetWindow.sync(msg);

  // update ship fitting and cargo window
  engineSack.shipFittingWindow.syncProperty(msg);
}

function handleDockedUsersUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerDockedUsersUpdate;

  // update ship fitting and cargo window
  engineSack.shipFittingWindow.syncDockedUsers(msg);
}

function handleExperienceUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerExperienceUpdate;

  // fix missing lists
  if (!msg.modules) {
    msg.modules = [];
  }

  if (!msg.ships) {
    msg.ships = [];
  }

  // update experience sheet window
  engineSack.experienceSheetWindow.sync(msg);
}

function handleSchematicRunsUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerSchematicRunsUpdate;

  // fix missing lists
  if (!msg.runs) {
    msg.runs = [];
  }

  // update schematic runs window
  engineSack.schematicRunsWindow.sync(msg);
}

function handleApplicationsUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerApplicationsUpdate;

  // fix missing lists
  if (!msg.applications) {
    msg.applications = [];
  }

  // update reputation sheet window
  engineSack.reputationSheetWindow.syncApplications(msg);
}

function handleMembersUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerMembersUpdate;

  // fix missing lists
  if (!msg.members) {
    msg.members = [];
  }

  // update reputation sheet window
  engineSack.reputationSheetWindow.syncMembers(msg);
}

function handleActionReportsPage(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerActionReportsPage;

  // fix missing lists
  if (!msg.logs) {
    msg.logs = [];
  }

  // fix null tickers
  for (const t of msg.logs) {
    if (!t.ticker) {
      t.ticker = '';
    }
  }

  // update action reports window
  engineSack.actionReportsWindow.setPageData(msg);
}

function handleActionReportDetails(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body);

  // export as pretty json
  const pretty = JSON.stringify(msg, null, 2);

  // copy to clipboard
  copyToClipboard(pretty, 'action report data copied to clipboard');
}

// clears the screen
function gfxBlank() {
  engineSack.ctx.fillStyle = 'pink';
  engineSack.ctx.fillRect(0, 0, engineSack.gfx.width, engineSack.gfx.height);
}

// indicator that the player is docked
function gfxDockOverlay() {
  // draw docked background
  engineSack.ctx.fillStyle = 'black';
  engineSack.ctx.strokeStyle = 'black';
  engineSack.ctx.fillRect(
    engineSack.gfx.width / 2 - 100,
    engineSack.gfx.height / 2 - 25,
    200,
    50,
  );

  // draw docked text
  engineSack.ctx.fillStyle = 'gray';
  engineSack.ctx.font = '30px FiraCode-Regular';
  engineSack.ctx.fillText(
    'Docked',
    engineSack.gfx.width / 2 - 100,
    engineSack.gfx.height / 2,
  );
}

// draws the backplate for the current system
function gfxBackplate() {
  if (
    !engineSack.player.currentSystem.backplateImg ||
    !engineSack.player.currentSystem.backplateValid
  ) {
    // render backplate
    engineSack.backplateRenderer.render({
      renderPointStars: true,
      renderStars: true,
      renderSun: false,
      renderNebulae: true,
      shortScale: false,
      seed: engineSack.player.currentSystem.id, // quick way to get a different plate for each system
    });

    // get data url and convert to image
    engineSack.player.currentSystem.backplateImg = new Image();
    engineSack.player.currentSystem.backplateImg.src =
      engineSack.backplateCanvas.toDataURL('image/png');

    // force refresh of backplate canvas
    const w = engineSack.backplateCanvas.width;
    const h = engineSack.backplateCanvas.height;

    engineSack.backplateCanvas.width = 0;
    engineSack.backplateCanvas.height = 0;

    setTimeout(() => {
      engineSack.backplateCanvas.width = w;
      engineSack.backplateCanvas.height = h;
    }, 0);

    // mark as valid
    engineSack.player.currentSystem.backplateValid = true;
  }

  // draw system backplate
  engineSack.ctx.drawImage(engineSack.player.currentSystem.backplateImg, 0, 0);
}

function clientLoop() {
  // periodic update
  periodicUpdate();

  // render
  try {
    clientRender();
  } catch (error) {
    console.log('caught error in client render');
    console.log(error);
  }

  // check if connection has been lost
  if (engineSack.wsSvc.isStale() && !engineSack.reloading) {
    engineSack.reloading = true;

    alert('Connection lost.');
    window.location.reload();
  }

  // clamp player's ship to interaction window as host
  engineSack.targetInteractionWindow.setHost(engineSack.player.currentShip);

  // calculate time since last frame
  engineSack.tpf = Date.now() - engineSack.lastFrameTime;

  // store frame time
  engineSack.lastFrameTime = Date.now();

  // queue next iteration
  setTimeout(clientLoop, 30 - engineSack.tpf);
}

function periodicUpdate() {
  // update target selection
  updateTargetSelection();

  // update position if docked
  if (!!engineSack.player.currentShip.dockedAtStationID) {
    // check outposts
    for (const st of engineSack.player.currentSystem.outposts) {
      if (st.id == engineSack.player.currentShip.dockedAtStationID) {
        engineSack.camera.x = st.x;
        engineSack.camera.y = st.y;
      }
    }

    // check stations
    for (const st of engineSack.player.currentSystem.stations) {
      if (st.id == engineSack.player.currentShip.dockedAtStationID) {
        engineSack.camera.x = st.x;
        engineSack.camera.y = st.y;
      }
    }
  }

  // update ui
  for (const w of engineSack.windows) {
    w.periodicUpdate();
  }

  // update module visual effects
  for (const ef of engineSack.player.currentSystem.moduleEffects) {
    ef.periodicUpdate();
  }

  // update point visual effects
  for (const ef of engineSack.player.currentSystem.pointEffects) {
    ef.periodicUpdate();
  }
}

function interpolate() {
  const start = Date.now();

  // calculate rolling q
  const q = calculateQ();

  // interate over ships
  for (const sh of engineSack.player.currentSystem.ships) {
    // get average delta
    const delta = sh.getAverageSyncDelta();

    // get frame velocity
    const fVx = delta.deltaX / q;
    const fVy = delta.deltaY / q;

    // adjust position
    sh.x += fVx;
    sh.y += fVy;
  }

  // calculate tpi
  const done = Date.now();
  engineSack.tpi = done - start;

  // clamp tpi
  if (engineSack.tpi < 5) {
    engineSack.tpi = 5;
  }

  // queue next iteration
  setTimeout(interpolate, engineSack.tpi);
}

function calculateQ(): number {
  // average sync tail
  let avgTps = 0;

  for (const qx of engineSack.tpsTail) {
    avgTps += qx;
  }

  avgTps /= engineSack.tpsTail.length;

  // get scale
  const q = Math.round(Math.max(avgTps / engineSack.tpi, 40 / engineSack.tpi));

  // return result
  return q;
}

function updateTargetSelection() {
  // wrecks
  for (const sm of engineSack.player.currentSystem.wrecks) {
    // current ship target check if undocked
    if (!engineSack.player.currentShip.dockedAtStationID) {
      if (
        sm.id === engineSack.player.currentTargetID &&
        engineSack.player.currentTargetType === TargetType.Wreck
      ) {
        // mark as targeted
        sm.isTargeted = true;
        engineSack.targetInteractionWindow.setTarget(sm, TargetType.Wreck);
      } else {
        // mark as untargeted
        sm.isTargeted = false;
      }
    }
  }

  // stars
  for (const sm of engineSack.player.currentSystem.stars) {
    // current ship target check if undocked
    if (!engineSack.player.currentShip.dockedAtStationID) {
      if (
        sm.id === engineSack.player.currentTargetID &&
        engineSack.player.currentTargetType === TargetType.Star
      ) {
        // mark as targeted
        sm.isTargeted = true;
        engineSack.targetInteractionWindow.setTarget(sm, TargetType.Star);
      } else {
        // mark as untargeted
        sm.isTargeted = false;
      }
    }
  }

  // planets
  for (const sm of engineSack.player.currentSystem.planets) {
    // current ship target check if undocked
    if (!engineSack.player.currentShip.dockedAtStationID) {
      if (
        sm.id === engineSack.player.currentTargetID &&
        engineSack.player.currentTargetType === TargetType.Planet
      ) {
        // mark as targeted
        sm.isTargeted = true;
        engineSack.targetInteractionWindow.setTarget(sm, TargetType.Planet);
      } else {
        // mark as untargeted
        sm.isTargeted = false;
      }
    }
  }

  // jumpholes
  for (const sm of engineSack.player.currentSystem.jumpholes) {
    // current ship target check if undocked
    if (!engineSack.player.currentShip.dockedAtStationID) {
      if (
        sm.id === engineSack.player.currentTargetID &&
        engineSack.player.currentTargetType === TargetType.Jumphole
      ) {
        // mark as targeted
        sm.isTargeted = true;
        engineSack.targetInteractionWindow.setTarget(sm, TargetType.Jumphole);
      } else {
        // mark as untargeted
        sm.isTargeted = false;
      }
    }
  }

  // asteroids
  for (const sm of engineSack.player.currentSystem.asteroids) {
    // current ship target check if undocked
    if (!engineSack.player.currentShip.dockedAtStationID) {
      if (
        sm.id === engineSack.player.currentTargetID &&
        engineSack.player.currentTargetType === TargetType.Asteroid
      ) {
        // mark as targeted
        sm.isTargeted = true;
        engineSack.targetInteractionWindow.setTarget(sm, TargetType.Asteroid);
      } else {
        // mark as untargeted
        sm.isTargeted = false;
      }
    }
  }

  // stations
  for (const sm of engineSack.player.currentSystem.stations) {
    // current ship target check if undocked
    if (!engineSack.player.currentShip.dockedAtStationID) {
      if (
        sm.id === engineSack.player.currentTargetID &&
        engineSack.player.currentTargetType === TargetType.Station
      ) {
        // mark as targeted
        sm.isTargeted = true;
        engineSack.targetInteractionWindow.setTarget(sm, TargetType.Station);
      } else {
        // mark as untargeted
        sm.isTargeted = false;
      }
    } else {
      if (engineSack.player.currentShip.dockedAtStationID == sm.id) {
        // mark as targeted
        sm.isTargeted = true;
        engineSack.targetInteractionWindow.setTarget(sm, TargetType.Station);
      }
    }
  }

  // outposts
  for (const sm of engineSack.player.currentSystem.outposts) {
    // current ship target check if undocked
    if (!engineSack.player.currentShip.dockedAtStationID) {
      if (
        sm.id === engineSack.player.currentTargetID &&
        engineSack.player.currentTargetType === TargetType.Outpost
      ) {
        // mark as targeted
        sm.isTargeted = true;
        engineSack.targetInteractionWindow.setTarget(sm, TargetType.Outpost);
      } else {
        // mark as untargeted
        sm.isTargeted = false;
      }
    } else {
      if (engineSack.player.currentShip.dockedAtStationID == sm.id) {
        // mark as targeted
        sm.isTargeted = true;
        engineSack.targetInteractionWindow.setTarget(sm, TargetType.Outpost);
      }
    }
  }

  // ships
  for (const sm of engineSack.player.currentSystem.ships) {
    // current ship target check if undocked
    if (!engineSack.player.currentShip.dockedAtStationID) {
      if (
        sm.id === engineSack.player.currentTargetID &&
        engineSack.player.currentTargetType === TargetType.Ship
      ) {
        // mark as targeted
        sm.isTargeted = true;
        engineSack.targetInteractionWindow.setTarget(sm, TargetType.Ship);
      } else {
        // mark as untargeted
        sm.isTargeted = false;
      }
    }
  }
}

function clientRender() {
  // alternate frame flag
  engineSack.alternateFrame = !engineSack.alternateFrame;

  // get camera look from interaction window
  const look = engineSack.targetInteractionWindow.getCameraLook();

  if (look[0]) {
    // center camera on look
    engineSack.camera.x = look[0].x;
    engineSack.camera.y = look[0].y;
  } else {
    // center camera on player ship
    for (const sh of engineSack.player.currentSystem.ships) {
      if (sh.id == engineSack.player.currentShip.id) {
        engineSack.camera.x = sh.x;
        engineSack.camera.y = sh.y;

        break;
      }
    }
  }

  // blank screen
  gfxBlank();

  // draw system backplate
  gfxBackplate();

  // draw stars
  for (const st of engineSack.player.currentSystem.stars) {
    st.render(engineSack.ctx, engineSack.camera);
  }

  // draw planets
  for (const p of engineSack.player.currentSystem.planets) {
    p.render(engineSack.ctx, engineSack.camera);
  }

  // draw artifacts
  for (const af of engineSack.player.currentSystem.artifacts) {
    af.render(engineSack.ctx, engineSack.camera);
  }

  // draw jumpholes
  for (const j of engineSack.player.currentSystem.jumpholes) {
    j.render(engineSack.ctx, engineSack.camera);
  }

  // draw npc stations
  for (const st of engineSack.player.currentSystem.stations) {
    st.render(engineSack.ctx, engineSack.camera);
  }

  // draw player-owned stations
  const keepOutposts: Outpost[] = [];
  for (const st of engineSack.player.currentSystem.outposts) {
    // only draw outposts we've recently seen
    if (st.lastSeen > engineSack.lastSyncTime - (engineSack.tpf - 2)) {
      st.render(engineSack.ctx, engineSack.camera);
      keepOutposts.push(st);
    }
  }

  // draw asteroids
  for (const p of engineSack.player.currentSystem.asteroids) {
    p.render(engineSack.ctx, engineSack.camera);
  }

  // draw wrecks
  const keepWrecks: Wreck[] = [];
  for (const sh of engineSack.player.currentSystem.wrecks) {
    // only draw wrecks we've "recently" seen
    if (Date.now() - sh.lastSeen <= 5000) {
      sh.render(engineSack.ctx, engineSack.camera);
      keepWrecks.push(sh);
    }
  }

  // draw ships
  const keepShips: Ship[] = [];
  for (const sh of engineSack.player.currentSystem.ships) {
    // only draw ships we've recently seen
    if (sh.lastSeen > engineSack.lastSyncTime - (engineSack.tpf - 2)) {
      sh.render(engineSack.ctx, engineSack.camera);
      keepShips.push(sh);
    }
  }

  // draw missiles
  const keepMissiles: Missile[] = [];
  for (const sh of engineSack.player.currentSystem.missiles) {
    // only draw missiles we've recently seen
    if (sh.lastSeen > engineSack.lastSyncTime - (engineSack.tpf - 2)) {
      sh.render(engineSack.ctx, engineSack.camera);
      keepMissiles.push(sh);
    }
  }

  // draw module visual effects
  for (const ef of engineSack.player.currentSystem.moduleEffects) {
    ef.render(engineSack.ctx, engineSack.camera);
  }

  // draw point visual effects
  for (const ef of engineSack.player.currentSystem.pointEffects) {
    ef.render(engineSack.ctx, engineSack.camera);
  }

  // keep only ships / missiles / wrecks / outposts that were drawable in-memory
  engineSack.player.currentSystem.ships = keepShips;
  engineSack.player.currentSystem.missiles = keepMissiles;
  engineSack.player.currentSystem.wrecks = keepWrecks;
  engineSack.player.currentSystem.outposts = keepOutposts;

  // draw overlay if docked
  if (!!engineSack.player.currentShip.dockedAtStationID) {
    gfxDockOverlay();
  }

  // draw ui windows (from bottom to top)
  for (const w of engineSack.windows.slice().reverse()) {
    if (w.isHidden()) {
      continue;
    }

    let bmp: ImageBitmap = null;

    if (engineSack.alternateFrame) {
      // render new frame
      bmp = w.render();
    } else {
      // reshow last frame
      bmp = w.getLastRender();
    }

    if (bmp != null) {
      engineSack.ctx.drawImage(bmp, w.getX(), w.getY());
    }
  }
}

function handleClick(x: number, y: number) {
  // check to see if we're clicking on any windows
  for (const w of engineSack.windows) {
    if (w.containsPoint(x, y)) {
      // allow window to handle click
      w.handleClick(x, y);
      return;
    }
  }

  // shift down means targeting click instead of nav click
  if (Date.now() - engineSack.lastShiftDown < 200) {
    // check to see if we're clicking on any ships
    for (const sh of engineSack.player.currentSystem.ships) {
      // skip if player ship
      if (sh.id === engineSack.player.currentShip.id) {
        continue;
      }

      // project coordinates to screen
      const sX = engineSack.camera.projectX(sh.x);
      const sY = engineSack.camera.projectY(sh.y);
      const sR = engineSack.camera.projectR(sh.radius);

      // check for intersection
      const m = magnitude(x, y, sX, sY);
      if (m < sR) {
        // set as target on client
        engineSack.player.currentTargetID = sh.id;
        engineSack.player.currentTargetType = TargetType.Ship;

        // store target on overview window as well
        engineSack.overviewWindow.selectedItemID = sh.id;
        engineSack.overviewWindow.selectedItemType = TargetType.Ship;

        return;
      }
    }

    // check to see if we're clicking on any stations
    for (const st of engineSack.player.currentSystem.stations) {
      // project coordinates to screen
      const sX = engineSack.camera.projectX(st.x);
      const sY = engineSack.camera.projectY(st.y);
      const sR = engineSack.camera.projectR(st.radius);

      // check for intersection
      const m = magnitude(x, y, sX, sY);
      if (m < sR) {
        // set as target on client
        engineSack.player.currentTargetID = st.id;
        engineSack.player.currentTargetType = TargetType.Station;

        // store target on overview window as well
        engineSack.overviewWindow.selectedItemID = st.id;
        engineSack.overviewWindow.selectedItemType = TargetType.Station;

        return;
      }
    }

    // check to see if we're clicking on any outposts
    for (const st of engineSack.player.currentSystem.outposts) {
      // project coordinates to screen
      const sX = engineSack.camera.projectX(st.x);
      const sY = engineSack.camera.projectY(st.y);
      const sR = engineSack.camera.projectR(st.radius);

      // check for intersection
      const m = magnitude(x, y, sX, sY);
      if (m < sR) {
        // set as target on client
        engineSack.player.currentTargetID = st.id;
        engineSack.player.currentTargetType = TargetType.Outpost;

        // store target on overview window as well
        engineSack.overviewWindow.selectedItemID = st.id;
        engineSack.overviewWindow.selectedItemType = TargetType.Outpost;

        return;
      }
    }
  } else {
    // issue a nav order for that location in space
    const b = new ClientNavClick();

    // calculate cursor vector
    b.dT = angleBetween(
      engineSack.gfx.width / 2,
      engineSack.gfx.height / 2,
      x,
      y,
    );
    b.m =
      magnitude(engineSack.gfx.width / 2, engineSack.gfx.height / 2, x, y) /
      // half way across the shortest part of the screen is a full speed request
      Math.min(engineSack.gfx.width / 4, engineSack.gfx.height / 4);

    // send nav click request
    b.sid = engineSack.player.sid;
    engineSack.wsSvc.sendMessage(MessageTypes.NavClick, b);
  }
}

function handleMouseMove(x: number, y: number) {
  // handle dragging window
  for (const w of engineSack.windows) {
    if (w.isDragging()) {
      // move window following cursor
      w.setX(x);
      w.setY(y);

      // fix boundary crossing
      if (w.getX() + w.getWidth() > engineSack.gfx.width) {
        w.setX(engineSack.gfx.width - w.getWidth());
      }

      if (
        w.getY() + (w.getHeight() + GDIStyle.windowHandleHeight) >
        engineSack.gfx.height
      ) {
        w.setY(
          engineSack.gfx.height - (w.getHeight() + GDIStyle.windowHandleHeight),
        );
      }

      // move dragged window to top
      engineSack.windows = [
        w,
        ...engineSack.windows.filter((item) => item !== w),
      ];

      // only allow one window to drag at a time
      return;
    } else {
      // pass event to window
      w.handleMouseMove(x, y);
    }
  }

  // cache last known mouse position
  engineSack.mouseX = x;
  engineSack.mouseY = y;
}

function handleScroll(dY: number) {
  // check to see if we're scrolling in any windows
  const x = engineSack.mouseX;
  const y = engineSack.mouseY;

  for (const w of engineSack.windows) {
    if (w.containsPoint(x, y)) {
      // allow window to handle scroll
      w.handleScroll(x, y, dY);
      return;
    }
  }

  // zoom camera
  if (dY < 0) {
    // zoom in
    engineSack.camera.zoom *= 1.1;
  } else {
    // zoom out
    engineSack.camera.zoom *= 0.9;
  }
}

function handleKeydown(key: string) {
  // check for shift key
  if (key === 'Shift') {
    engineSack.lastShiftDown = Date.now();
  }

  // check to see if we're typing in any windows
  const x = engineSack.mouseX;
  const y = engineSack.mouseY;

  for (const w of engineSack.windows) {
    if (w.containsPoint(x, y)) {
      // allow window to handle key press
      w.handleKeyDown(x, y, key);
      return;
    }
  }
}

function handleWindowResize() {
  // get current window size
  const clientWidth = document.documentElement.clientWidth;
  const clientHeight = document.documentElement.clientHeight;

  // resize canvases
  engineSack.gfx.width = clientWidth;
  engineSack.gfx.height = clientHeight;

  engineSack.backplateCanvas.width = clientWidth;
  engineSack.backplateCanvas.height = clientHeight;

  // resize camera viewport
  engineSack.camera.resizeViewport(clientWidth, clientHeight);

  // resize window manager
  engineSack.windowManager.resize(clientHeight);

  // make sure all windows are within the new viewport
  for (const w of engineSack.windows) {
    while (w.getX() + w.getWidth() + 1 > clientWidth) {
      w.setX(w.getX() - 1);
    }

    while (w.getY() + w.getHeight() + 1 > clientHeight) {
      w.setY(w.getY() - 1);
    }

    if (w.getX() < 0) {
      w.setX(0);
    }

    if (w.getY() < 0) {
      w.setY(0);
    }
  }

  // mark backplate as dirty
  engineSack.player.currentSystem.backplateValid = false;
}

function copyToClipboard(s: string, msg: string) {
  // copy text to clipboard
  ClipboardJS.copy(s);

  // move push info window to top
  engineSack.windows = [
    engineSack.pushInfoWindow,
    ...engineSack.windows.filter((item) => item !== engineSack.pushInfoWindow),
  ];

  // show the push info window
  engineSack.pushInfoWindow.setText(msg);
  engineSack.pushInfoWindow.setHidden(false);
}
