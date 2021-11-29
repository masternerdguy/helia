import { GDIWindow } from '../base/gdiWindow';
import { FontSize, GDIStyle } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { WsService } from '../../ws.service';
import { GDIInput } from '../components/gdiInput';
import { WSSystemChatMessage } from '../../wsModels/entities/wsSystemChatMessage';
import { ClientPostSystemChatMessage } from '../../wsModels/bodies/postSystemChatMessage';
import { MessageTypes } from '../../wsModels/gameMessage';
import { heliaDate, printHeliaDate } from '../../engineMath';

export class SystemChatWindow extends GDIWindow {
  private chatList = new GDIList();
  private knownMessages: KnownChatMessage[];

  // inputs
  private chatInput: GDIInput = new GDIInput();

  // ws service
  private wsSvc: WsService;

  initialize() {
    // set dimensions
    this.setWidth(400);
    this.setHeight(300);

    // initialize
    super.initialize();
    this.knownMessages = [];
  }

  pack() {
    this.setTitle('System Chat');

    // setup chat input
    const fontSize = GDIStyle.getUnderlyingFontSize(FontSize.large);
    this.chatInput.setWidth(400);
    this.chatInput.setHeight(Math.round(fontSize + 1.5));
    this.chatInput.setX(0);
    this.chatInput.setY(this.getHeight() - this.chatInput.getHeight());
    this.chatInput.setFont(FontSize.large);
    this.chatInput.initialize();

    this.chatInput.setOnReturn((txt: string) => {
      // send chat message
      const b = new ClientPostSystemChatMessage();
      b.sid = this.wsSvc.sid;
      b.message = txt;

      this.wsSvc.sendMessage(MessageTypes.PostSystemChatMessage, b);

      // clear input
      this.chatInput.setText('');
    });

    // setup chat message list
    this.chatList.setWidth(400);
    this.chatList.setHeight(this.getHeight() - this.chatInput.getHeight());
    this.chatList.initialize();
    this.chatList.setX(0);
    this.chatList.setY(0);
    this.chatList.setFont(FontSize.large);

    // pack
    this.addComponent(this.chatList);
    this.addComponent(this.chatInput);
  }

  periodicUpdate() {
    //
  }

  setWsService(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  sync(newMessages: WSSystemChatMessage[]) {
    if (newMessages.length == 0) {
      return;
    }

    // store messages
    for (const m of newMessages) {
      this.knownMessages.push({
        message: m,
        received: Date.now(),
      });
    }

    // update chat list
    this.chatList.setItems([]);
    const rows: ChatViewRow[] = [];

    let lastSenderId = '';
    let lastMessageReceived = 0;

    for (const m of this.knownMessages) {
      if (lastSenderId != m.message.senderId) {
        lastSenderId = m.message.senderId;

        // append sender info
        rows.push({
          listString: () => `[${m.message.senderName}]`,
        });
      }

      // parse timestamp as helia date
      const date = heliaDate(new Date(m.received));
      let broken: any[];

      if ((m.received - lastMessageReceived) / 1000 > 60) {
        // include timestamp
        const stamped = `${printHeliaDate(date)}\n${m.message.message}`;

        // break text to fit
        broken = this.chatList.breakText(stamped);
      } else {
        // break text to fit
        broken = this.chatList.breakText(m.message.message);
      }

      // append message text
      for (const b of broken) {
        rows.push({
          listString: () => b.text,
        });
      }

      // update timestamp tracker
      lastMessageReceived = m.received;
    }

    // store and scroll to bottom
    this.chatList.setItems(rows);
    this.chatList.setScroll(Number.MAX_SAFE_INTEGER);
  }
}

class ChatViewRow {
  listString: () => string;
}

class KnownChatMessage {
  message: WSSystemChatMessage;
  received: number;
}
