import { GDIWindow } from '../base/gdiWindow';
import { TargetType } from '../../engineModels/player';
import { GDIButton } from '../components/gdiButton';
import { FontSize } from '../base/gdiStyle';
import { WsService } from '../../ws.service';
import { ClientGoto } from '../../wsModels/bodies/goto';
import { MessageTypes } from '../../wsModels/gameMessage';

export class TargetInteractionWindow extends GDIWindow {
    private target: any;
    private targetType: TargetType;

    private gotoBtn = new GDIButton();
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

        this.addComponent(this.gotoBtn);
    }

    periodicUpdate() {
        if (this.target !== undefined) {
            this.gotoBtn.setEnabled(true);
        } else {
            this.gotoBtn.setEnabled(false);
        }
    }

    setTarget(target: any, targetType: TargetType) {
        this.target = target;
        this.targetType = targetType;
    }

    setWsSvc(wsSvc: WsService) {
        this.wsSvc = wsSvc;
    }
}
