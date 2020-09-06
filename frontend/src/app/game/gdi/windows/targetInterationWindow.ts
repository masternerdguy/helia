import { GDIWindow } from '../base/gdiWindow';
import { TargetType } from '../../engineModels/player';
import { GDIButton } from '../components/gdiButton';
import { FontSize } from '../base/gdiStyle';
import { WsService } from '../../ws.service';
import { ClientGoto } from '../../wsModels/bodies/goto';
import { MessageTypes } from '../../wsModels/gameMessage';
import { ClientOrbit } from '../../wsModels/bodies/orbit';
import { ClientDock } from '../../wsModels/bodies/dock';
import { Ship } from '../../engineModels/ship';
import { ClientUndock } from '../../wsModels/bodies/undock';

export class TargetInteractionWindow extends GDIWindow {
    private target: any;
    private targetType: TargetType;
    private host: Ship;

    private gotoBtn = new GDIButton();
    private orbitBtn = new GDIButton();
    private dockBtn = new GDIButton();
    private wsSvc: WsService;

    initialize() {
        // set dimensions
        this.setWidth(200);
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

        this.addComponent(this.gotoBtn);
        this.addComponent(this.orbitBtn);
        this.addComponent(this.dockBtn);
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
}
