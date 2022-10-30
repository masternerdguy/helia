import { GDIWindow } from '../base/gdiWindow';
import { TargetType } from '../../engineModels/player';
import { GDIButton } from '../components/gdiButton';
import { FontSize, GDIStyle } from '../base/gdiStyle';
import { WsService } from '../../ws.service';
import { ClientGoto } from '../../wsModels/bodies/goto';
import { MessageTypes } from '../../wsModels/gameMessage';
import { ClientOrbit } from '../../wsModels/bodies/orbit';
import { ClientDock } from '../../wsModels/bodies/dock';
import { Ship } from '../../engineModels/ship';
import { ClientUndock } from '../../wsModels/bodies/undock';
import { GDIBar } from '../components/gdiBar';

export class TargetInteractionWindow extends GDIWindow {
  private target: any;
  private targetType: TargetType;
  private host: Ship;

  private gotoBtn = new GDIButton();
  private orbitBtn = new GDIButton();
  private dockBtn = new GDIButton();
  private lookBtn = new GDIButton();
  private wsSvc: WsService;

  private tgtShieldBar = new GDIBar();
  private tgtArmorBar = new GDIBar();
  private tgtHullBar = new GDIBar();

  private cameraLook: any;

  initialize() {
    // set dimensions
    this.setWidth(370);
    this.setHeight(40);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Interaction');

    // goto button
    this.gotoBtn.setWidth(30);
    this.gotoBtn.setHeight(30);
    this.gotoBtn.initialize();

    this.gotoBtn.setX(5);
    this.gotoBtn.setY(5);

    this.gotoBtn.setFont(FontSize.giant);
    this.gotoBtn.setText('⤒');

    this.gotoBtn.setOnClick((x, y) => {
      // issue a goto order for selected target
      const b = new ClientGoto();
      b.sid = this.wsSvc.sid;
      b.targetId = this.target.id;
      b.type = this.targetType;

      this.wsSvc.sendMessage(MessageTypes.Goto, b);
    });

    // orbit button
    this.orbitBtn.setWidth(30);
    this.orbitBtn.setHeight(30);
    this.orbitBtn.initialize();

    this.orbitBtn.setX(40);
    this.orbitBtn.setY(5);

    this.orbitBtn.setFont(FontSize.giant);
    this.orbitBtn.setText('⥁');

    this.orbitBtn.setOnClick((x, y) => {
      // issue an orbit order for selected target
      const b = new ClientOrbit();
      b.sid = this.wsSvc.sid;
      b.targetId = this.target.id;
      b.type = this.targetType;

      this.wsSvc.sendMessage(MessageTypes.Orbit, b);
    });

    // dock button
    this.dockBtn.setWidth(30);
    this.dockBtn.setHeight(30);
    this.dockBtn.initialize();

    this.dockBtn.setX(75);
    this.dockBtn.setY(5);

    this.dockBtn.setFont(FontSize.giant);
    this.dockBtn.setText('⇴');

    this.dockBtn.setOnClick((x, y) => {
      if (!this.host.dockedAtStationID) {
        // issue a dock order for selected target
        const b = new ClientDock();
        b.sid = this.wsSvc.sid;
        b.targetId = this.target.id;
        b.type = this.targetType;

        this.wsSvc.sendMessage(MessageTypes.Dock, b);
      } else {
        // issue an undock order
        const b = new ClientUndock();
        b.sid = this.wsSvc.sid;

        this.wsSvc.sendMessage(MessageTypes.Undock, b);
      }
    });

    // look button
    this.lookBtn.setWidth(30);
    this.lookBtn.setHeight(30);
    this.lookBtn.initialize();

    this.lookBtn.setX(110);
    this.lookBtn.setY(5);

    this.lookBtn.setFont(FontSize.giant);
    this.lookBtn.setText('Ꙫ');

    this.lookBtn.setOnClick((x, y) => {
      if (!this.cameraLook) {
        // reparent camera to selected target
        this.cameraLook = this.target;
      } else {
        // reparent camera to player ship
        this.cameraLook = this.host;
      }
    });

    this.addComponent(this.gotoBtn);
    this.addComponent(this.orbitBtn);
    this.addComponent(this.dockBtn);

    // setup shield bar
    this.tgtShieldBar.setWidth(110);
    this.tgtShieldBar.setHeight(13.333333);
    this.tgtShieldBar.initialize();

    this.tgtShieldBar.setX(260);
    this.tgtShieldBar.setY(0);
    this.tgtShieldBar.setPercentage(0);

    this.tgtShieldBar.setFont(FontSize.small);
    this.tgtShieldBar.setShowPercentage(true);
    this.tgtShieldBar.setColor(GDIStyle.shieldBarColor);

    // setup armor bar
    this.tgtArmorBar.setWidth(110);
    this.tgtArmorBar.setHeight(13.333333);
    this.tgtArmorBar.initialize();

    this.tgtArmorBar.setX(260);
    this.tgtArmorBar.setY(13.333333);
    this.tgtArmorBar.setPercentage(0);

    this.tgtArmorBar.setFont(FontSize.small);
    this.tgtArmorBar.setShowPercentage(true);
    this.tgtArmorBar.setColor(GDIStyle.armorBarColor);

    // setup hull bar
    this.tgtHullBar.setWidth(110);
    this.tgtHullBar.setHeight(13.333333);
    this.tgtHullBar.initialize();

    this.tgtHullBar.setX(260);
    this.tgtHullBar.setY(26.666666);
    this.tgtHullBar.setPercentage(0);

    this.tgtHullBar.setFont(FontSize.small);
    this.tgtHullBar.setShowPercentage(true);
    this.tgtHullBar.setColor(GDIStyle.hullBarColor);

    this.addComponent(this.tgtShieldBar);
    this.addComponent(this.tgtArmorBar);
    this.addComponent(this.tgtHullBar);
  }

  periodicUpdate() {
    if (!this.host) {
      return;
    }

    if (!this.host.dockedAtStationID) {
      // player is in space
      if (this.target !== undefined) {
        // global buttons
        this.gotoBtn.setEnabled(true);
        this.orbitBtn.setEnabled(true);

        // station-specific buttons
        if (this.targetType === TargetType.Station) {
          this.dockBtn.setEnabled(true);

          // use dock icon
          this.dockBtn.setText('⇴');
        } else {
          this.dockBtn.setEnabled(false);
        }
      } else {
        // nothing selected
        this.gotoBtn.setEnabled(false);
        this.orbitBtn.setEnabled(false);
        this.dockBtn.setEnabled(false);
      }
    } else {
      // player is docked
      this.gotoBtn.setEnabled(false);
      this.orbitBtn.setEnabled(false);
      this.dockBtn.setEnabled(true);

      // use undock icon
      this.dockBtn.setText('⬰');
    }

    // update health bars
    if (this.targetType === TargetType.Ship) {
      const ct = this.target as Ship;
      this.showHealthBarInfo(ct?.shieldP, ct?.armorP, ct?.hullP);
    } else {
      this.hideHealthBarInfo();
    }
  }

  setTarget(target: any, targetType: TargetType) {
    this.target = target;
    this.targetType = targetType;
  }

  setHost(host: Ship) {
    this.host = host;
  }

  setWsSvc(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  getCameraLook() {
    return this.cameraLook;
  }

  private showHealthBarInfo(shieldP: number, armorP: number, hullP: number) {
    this.tgtShieldBar.setShowPercentage(true);
    this.tgtArmorBar.setShowPercentage(true);
    this.tgtHullBar.setShowPercentage(true);

    this.tgtShieldBar.setPercentage(shieldP ?? 0);
    this.tgtArmorBar.setPercentage(armorP ?? 0);
    this.tgtHullBar.setPercentage(hullP ?? 0);
  }

  private hideHealthBarInfo() {
    this.tgtShieldBar.setShowPercentage(false);
    this.tgtArmorBar.setShowPercentage(false);
    this.tgtHullBar.setShowPercentage(false);

    this.tgtShieldBar.setPercentage(0);
    this.tgtArmorBar.setPercentage(0);
    this.tgtHullBar.setPercentage(0);
  }
}
