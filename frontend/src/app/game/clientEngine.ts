import { WsService } from './ws.service';
import { ServerJoinBody } from './wsModels/bodies/join';
import { GameMessage, MessageTypes } from './wsModels/gameMessage';
import { Player } from './engineModels/player';
import { System } from './engineModels/system';
import { Ship } from './engineModels/ship';
import { Camera } from './engineModels/camera';
import { ServerGlobalUpdateBody } from './wsModels/bodies/globalUpdate';
import { Backplate } from './procedural/backplate/backplate';
import { ClientNavClick } from './wsModels/bodies/navClick';
import { angleBetween, magnitude } from './engineMath';
import { Star } from './engineModels/star';
import { Planet } from './engineModels/planet';
import { TestWindow } from './gdi/windows/testWindow';
import { GDIWindow } from './gdi/base/gdiWindow';
import { Station } from './engineModels/station';
import { GDIStyle } from './gdi/base/gdiStyle';
import { Jumphole } from './engineModels/jumphole';
import { ShipStatusWindow } from './gdi/windows/shipStatusWindow';
import { ServerCurrentShipUpdate } from './wsModels/bodies/currentShipUpdate';

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

  windows: GDIWindow[];

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

export function clientStart(wsService: WsService, gameCanvas: HTMLCanvasElement, backCanvas: HTMLCanvasElement, sid: string) {
  // initialize
  engineSack.player = new Player();
  engineSack.camera = new Camera(gameCanvas.width, gameCanvas.height, 1);
  engineSack.backplateRenderer = new Backplate(backCanvas);

  // initialize ui
  engineSack.shipStatusWindow = new ShipStatusWindow();
  engineSack.shipStatusWindow.setX(100);
  engineSack.shipStatusWindow.setY(100);
  engineSack.shipStatusWindow.initialize();
  engineSack.shipStatusWindow.pack();

  // cache windows for simpler updating and rendering
  engineSack.windows = [engineSack.shipStatusWindow];

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
  engineSack.gfx.addEventListener('wheel', event => {
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
  } else {
    // fix empty arrays in incoming data
    if (!msg.planets || msg.planets == null) {
      msg.planets = [];
    }

    if (!msg.stations || msg.stations == null) {
      msg.stations = [];
    }

    // update system
    engineSack.player.currentSystem.id = msg.currentSystemInfo.id;
    engineSack.player.currentSystem.systemName = msg.currentSystemInfo.systemName;

    // update ships
    for (const sh of msg.ships) {
      let match = false;

      // find ship in memory
      for (const sm of engineSack.player.currentSystem.ships) {
        if (sh.id === sm.id) {
          match = true;

          // sync ship in memory
          sm.sync(sh);

          // is this the player ship?
          if (sm.id === engineSack.player.currentShip.id) {
            // update player ship cache
            engineSack.player.currentShip.sync(sh);

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

      // todo: handle npc station dying
    }
  }
}

function handleCurrentShipUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerCurrentShipUpdate;

  // update current ship cache
  engineSack.player.currentShip.sync(msg.currentShipInfo);

  // update status window
  engineSack.shipStatusWindow.setShip(engineSack.player.currentShip);
}

// clears the screen
function gfxBlank() {
  engineSack.ctx.fillStyle = 'pink';
  engineSack.ctx.fillRect(0, 0, engineSack.gfx.width, engineSack.gfx.height);
}

// draws the backplate for the current system
function gfxBackplate() {
  if (!engineSack.player.currentSystem.backplateImg || !engineSack.player.currentSystem.backplateValid) {
    // render backplate
    engineSack.backplateRenderer.render(
      {
        renderPointStars: true,
        renderStars: true,
        renderSun: false,
        renderNebulae: true,
        shortScale: false,
        seed: engineSack.player.currentSystem.id // quick way to get a different plate for each system
      }
    );

    // get data url and convert to image
    engineSack.player.currentSystem.backplateImg = new Image();
    engineSack.player.currentSystem.backplateImg.src = engineSack.backplateCanvas.toDataURL('image/png');

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

  // keep only ships that were drawable in-memory
  engineSack.player.currentSystem.ships = keepShips;

  // draw ui elements
  for (const w of engineSack.windows) {
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

  // clicked on empty space - issue a nav order for that location
  const b = new ClientNavClick();

  // calculate cursor vector
  b.dT = angleBetween(engineSack.gfx.width / 2, engineSack.gfx.height / 2, x, y);
  b.m = (magnitude(engineSack.gfx.width / 2, engineSack.gfx.height / 2, x, y))
    // half way across the shortest part of the screen is a full speed request
    / Math.min(engineSack.gfx.width / 4, engineSack.gfx.height / 4);

  // send nav click request
  b.sid = engineSack.player.sid;
  engineSack.wsSvc.sendMessage(MessageTypes.NavClick, b);
}

function handleMouseMove(x: number, y: number) {
  // handle dragging window
  for (const w of engineSack.windows) {
    if (w.isDragging()) {
      // move window following cursor
      w.setX(x);
      w.setY(y);

      // fix boundary crossing
      if ((w.getX() + w.getWidth()) > engineSack.gfx.width) {
        w.setX(engineSack.gfx.width - w.getWidth());
      }

      if ((w.getY() + (w.getHeight() + GDIStyle.windowHandleHeight)) > engineSack.gfx.height) {
        w.setY(engineSack.gfx.height - (w.getHeight() + GDIStyle.windowHandleHeight));
      }

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
