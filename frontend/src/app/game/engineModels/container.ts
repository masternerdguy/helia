import { WSContainer } from '../wsModels/entities/wsContainer';

export class Container extends WSContainer {
  constructor(s: WSContainer) {
    super();

    Object.assign(this, s);
  }
}
