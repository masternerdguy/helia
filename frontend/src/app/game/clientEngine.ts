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
  shipStatusWindow: ShipStatusWindow;
  targetInteractionWindow: TargetInteractionWindow;
  overviewWindow: OverviewWindow;
  shipFittingWindow: ShipFittingWindow;
  lastShiftDown: number;

  windows: GDIWindow[];
  windowManager: WindowManager;

  // client-server communication
  wsSvc: WsService;

  // tpf
  lastFrameTime: number;
  lastSyncTime: number;
  tpf: number;

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
  sid: string
) {
  // initialize
  engineSack.player = new Player();
  engineSack.camera = new Camera(gameCanvas.width, gameCanvas.height, 1);
  engineSack.backplateRenderer = new Backplate(backCanvas);

  // initialize window manager
  engineSack.windowManager = new WindowManager();
  engineSack.windowManager.preinit(gameCanvas.height);
  engineSack.windowManager.setX(0);
  engineSack.windowManager.setY(Number.NEGATIVE_INFINITY);
  engineSack.windowManager.initialize();

  // initialize ui windows
  engineSack.shipStatusWindow = new ShipStatusWindow();
  engineSack.shipStatusWindow.setX(engineSack.windowManager.getWidth());
  engineSack.shipStatusWindow.setY(0);
  engineSack.shipStatusWindow.initialize();
  engineSack.shipStatusWindow.pack();
  engineSack.shipStatusWindow.setWsService(wsService);
  engineSack.shipStatusWindow.setPlayer(engineSack.player);

  engineSack.shipFittingWindow = new ShipFittingWindow();
  engineSack.shipFittingWindow.setX(300);
  engineSack.shipFittingWindow.setY(300);
  engineSack.shipFittingWindow.initialize();
  engineSack.shipFittingWindow.pack();
  engineSack.shipFittingWindow.setWsService(wsService);
  engineSack.shipFittingWindow.setPlayer(engineSack.player);

  engineSack.targetInteractionWindow = new TargetInteractionWindow();
  engineSack.targetInteractionWindow.initialize();
  engineSack.targetInteractionWindow.setWsSvc(wsService);
  engineSack.targetInteractionWindow.setX(
    gameCanvas.width - engineSack.targetInteractionWindow.getWidth()
  );
  engineSack.targetInteractionWindow.setY(0);
  engineSack.targetInteractionWindow.pack();

  engineSack.overviewWindow = new OverviewWindow();
  engineSack.overviewWindow.setHeight(gameCanvas.height / 2);
  engineSack.overviewWindow.initialize();
  engineSack.overviewWindow.setX(
    gameCanvas.width - engineSack.overviewWindow.getWidth()
  );
  engineSack.overviewWindow.setY(
    engineSack.targetInteractionWindow.getY() +
      engineSack.targetInteractionWindow.getHeight() +
      GDIStyle.windowHandleHeight
  );
  engineSack.overviewWindow.pack();

  // link windows to window manager
  engineSack.windowManager.manageWindow(engineSack.overviewWindow, '☀');
  engineSack.windowManager.manageWindow(engineSack.shipStatusWindow, '☍');
  engineSack.windowManager.manageWindow(
    engineSack.targetInteractionWindow,
    '☉'
  );
  engineSack.windowManager.manageWindow(engineSack.shipFittingWindow, 'Ʌ');

  // cache windows for simpler updating and rendering
  engineSack.windows = [
    engineSack.shipStatusWindow,
    engineSack.targetInteractionWindow,
    engineSack.overviewWindow,
    engineSack.shipFittingWindow,
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

  // detect fullscreen loss
  const exitHandler = () => {
    window.location.href = '/auth/signin';
  };

  document.addEventListener('fullscreenchange', exitHandler, false);
  document.addEventListener('mozfullscreenchange', exitHandler, false);
  document.addEventListener('MSFullscreenChange', exitHandler, false);
  document.addEventListener('webkitfullscreenchange', exitHandler, false);

  // hide initially hidden windows
  engineSack.shipFittingWindow.setHidden(true);

  // start game loop
  engineSack.lastFrameTime = Date.now();
  engineSack.lastSyncTime = Date.now();
  engineSack.tpf = 0;

  setInterval(clientLoop, 20);
}

function handleGlobalUpdate(d: GameMessage) {
  // store sync time
  engineSack.lastSyncTime = Date.now();

  // parse body
  const msg = JSON.parse(d.body) as ServerGlobalUpdateBody;

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
  } else {
    // fix empty arrays in incoming data
    if (!msg.planets || msg.planets == null) {
      msg.planets = [];
    }

    if (!msg.stations || msg.stations == null) {
      msg.stations = [];
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

    if (!msg.newModuleEffects || msg.newModuleEffects == null) {
      msg.newModuleEffects = [];
    }

    if (!msg.newPointEffects || msg.newPointEffects == null) {
      msg.newPointEffects = [];
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

          // current ship target check if undocked
          if (!engineSack.player.currentShip.dockedAtStationID) {
            if (
              sm.id === engineSack.player.currentTargetID &&
              engineSack.player.currentTargetType === TargetType.Ship
            ) {
              sm.isTargeted = true;
              engineSack.targetInteractionWindow.setTarget(sm, TargetType.Ship);
            }
          }

          // is this the player ship?
          if (sm.id === engineSack.player.currentShip.id) {
            // flag as player's ship
            sm.isPlayer = true;

            // update player ship cache
            engineSack.player.currentShip.sync(sh);
            engineSack.player.currentShip.isPlayer = true;

            // update camera position to track player ship
            engineSack.camera.x = sm.x;
            engineSack.camera.y = sm.y;
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

    // update stars
    for (const st of msg.stars) {
      let match = false;

      // find star in memory
      for (const sm of engineSack.player.currentSystem.stars) {
        if (st.id === sm.id) {
          match = true;
          sm.isTargeted = false;

          // sync star in memory
          sm.sync(st);

          // current ship target check if undocked
          if (!engineSack.player.currentShip.dockedAtStationID) {
            if (
              sm.id === engineSack.player.currentTargetID &&
              engineSack.player.currentTargetType === TargetType.Star
            ) {
              sm.isTargeted = true;
              engineSack.targetInteractionWindow.setTarget(sm, TargetType.Star);
            }
          }

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
          sm.isTargeted = false;

          // sync planet in memory
          sm.sync(p);

          // current ship target check if undocked
          if (!engineSack.player.currentShip.dockedAtStationID) {
            if (
              sm.id === engineSack.player.currentTargetID &&
              engineSack.player.currentTargetType === TargetType.Planet
            ) {
              sm.isTargeted = true;
              engineSack.targetInteractionWindow.setTarget(
                sm,
                TargetType.Planet
              );
            }
          }

          // exit loop
          break;
        }
      }

      if (!match) {
        // add planet to memory
        engineSack.player.currentSystem.planets.push(new Planet(p));
      }
    }

    // update jumpholes
    for (const j of msg.jumpholes) {
      let match = false;

      // find jumphole in memory
      for (const sm of engineSack.player.currentSystem.jumpholes) {
        if (j.id === sm.id) {
          match = true;
          sm.isTargeted = false;

          // sync jumphole in memory
          sm.sync(j);

          // current ship target check if undocked
          if (!engineSack.player.currentShip.dockedAtStationID) {
            if (
              sm.id === engineSack.player.currentTargetID &&
              engineSack.player.currentTargetType === TargetType.Jumphole
            ) {
              sm.isTargeted = true;
              engineSack.targetInteractionWindow.setTarget(
                sm,
                TargetType.Jumphole
              );
            }
          }

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

          // current ship target check
          sm.isTargeted = false;

          if (
            sm.id === engineSack.player.currentTargetID &&
            engineSack.player.currentTargetType === TargetType.Station
          ) {
            sm.isTargeted = true;
            engineSack.targetInteractionWindow.setTarget(
              sm,
              TargetType.Station
            );
          }

          // exit loop
          break;
        }
      }

      if (!match) {
        // add station to memory
        engineSack.player.currentSystem.stations.push(new Station(p));
      }

      // todo: handle npc station dying
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

  console.log(vw);
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

  // update camera position to track player ship
  engineSack.camera.x = msg.currentShipInfo.x;
  engineSack.camera.y = msg.currentShipInfo.y;

  // dock check
  if (!!msg.currentShipInfo.dockedAtStationID) {
    engineSack.player.currentTargetID = msg.currentShipInfo.dockedAtStationID;
    engineSack.player.currentTargetType = TargetType.Station;

    // store target on overview window as well
    engineSack.overviewWindow.selectedItemID =
      msg.currentShipInfo.dockedAtStationID;
    engineSack.overviewWindow.selectedItemType = TargetType.Station;
  }

  // update status window
  engineSack.shipStatusWindow.setShip(engineSack.player.currentShip);

  // update fitting window
  engineSack.shipFittingWindow.setPlayer(engineSack.player);
  engineSack.shipFittingWindow.setWsService(engineSack.wsSvc);
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
    50
  );

  // draw docked text
  engineSack.ctx.fillStyle = 'gray';
  engineSack.ctx.font = '30px FiraCode-Regular';
  engineSack.ctx.fillText(
    'Docked',
    engineSack.gfx.width / 2 - 100,
    engineSack.gfx.height / 2
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
    engineSack.player.currentSystem.backplateImg.src = engineSack.backplateCanvas.toDataURL(
      'image/png'
    );

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
  clientRender();

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
}

function periodicUpdate() {
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

function clientRender() {
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

  // draw jumpholes
  for (const j of engineSack.player.currentSystem.jumpholes) {
    j.render(engineSack.ctx, engineSack.camera);
  }

  // draw npc stations
  for (const st of engineSack.player.currentSystem.stations) {
    st.render(engineSack.ctx, engineSack.camera);
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

  // draw module visual effects
  for (const ef of engineSack.player.currentSystem.moduleEffects) {
    ef.render(engineSack.ctx, engineSack.camera);
  }

  // draw point visual effects
  for (const ef of engineSack.player.currentSystem.pointEffects) {
    ef.render(engineSack.ctx, engineSack.camera);
  }

  // keep only ships that were drawable in-memory
  engineSack.player.currentSystem.ships = keepShips;

  // draw overlay if docked
  if (!!engineSack.player.currentShip.dockedAtStationID) {
    gfxDockOverlay();
  }

  // draw ui windows (from bottom to top)
  for (const w of engineSack.windows.slice().reverse()) {
    if (w.isHidden()) {
      continue;
    }

    const bmp = w.render();
    engineSack.ctx.drawImage(bmp, w.getX(), w.getY());
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
  } else {
    // issue a nav order for that location in space
    const b = new ClientNavClick();

    // calculate cursor vector
    b.dT = angleBetween(
      engineSack.gfx.width / 2,
      engineSack.gfx.height / 2,
      x,
      y
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
          engineSack.gfx.height - (w.getHeight() + GDIStyle.windowHandleHeight)
        );
      }

      // move dragged window to top
      engineSack.windows = [
        w,
        ...engineSack.windows.filter((item) => item !== w),
      ];

      // only allow one window to drag at a time
      return;
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
