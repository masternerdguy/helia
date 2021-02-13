import { WSContainerItem } from './wsContainer';

export class WSOpenSellOrdersUpdate {
  orders: WSOpenSellOrder[];
}

export class WSOpenSellOrder {
  id: string;
  stationId: string;
  itemId: string;
  sellerId: string;
  ask: number;
  createdAt: string;
  boughtAt: string;
  buyerId: string;
  item: WSContainerItem;
}
