import { FontSize } from '../base/gdiStyle';
import { GDIWindow } from '../base/gdiWindow';
import { GDIList } from '../components/gdiList';

export class PushErrorWindow extends GDIWindow {
  private textList = new GDIList();
  public testText =
`
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Pellentesque sit amet porttitor eget. At augue eget arcu dictum varius duis at consectetur. Odio pellentesque diam volutpat commodo. Nullam ac tortor vitae purus. Pharetra convallis posuere morbi leo urna. Purus sit amet luctus venenatis lectus magna. Eget duis at tellus at urna. Consectetur a erat nam at lectus urna duis convallis convallis. Est sit amet facilisis magna. Id interdum velit laoreet id donec ultrices.

Viverra maecenas accumsan lacus vel facilisis volutpat est. Ac felis donec et odio pellentesque diam volutpat commodo sed. Adipiscing bibendum est ultricies integer quis auctor. Arcu cursus euismod quis viverra. Sit amet mauris commodo quis imperdiet massa. Velit ut tortor pretium viverra. Faucibus a pellentesque sit amet porttitor eget dolor morbi. Justo nec ultrices dui sapien eget mi. Vehicula ipsum a arcu cursus vitae congue. Scelerisque fermentum dui faucibus in ornare.
    
Et leo duis ut diam. At augue eget arcu dictum varius duis at. Sit amet volutpat consequat mauris nunc congue nisi. Ornare arcu dui vivamus arcu felis bibendum ut tristique et. Ipsum dolor sit amet consectetur adipiscing elit ut aliquam. Scelerisque purus semper eget duis at tellus at urna condimentum. Ac tortor vitae purus faucibus ornare suspendisse sed nisi. Ipsum consequat nisl vel pretium lectus quam id leo. Rutrum tellus pellentesque eu tincidunt tortor aliquam nulla. Viverra adipiscing at in tellus integer feugiat scelerisque. Sit amet volutpat consequat mauris nunc congue nisi vitae. Pharetra pharetra massa massa ultricies mi quis hendrerit dolor magna. Interdum velit laoreet id donec ultrices tincidunt. Vitae purus faucibus ornare suspendisse. Elit ullamcorper dignissim cras tincidunt lobortis. Magna eget est lorem ipsum dolor sit amet. Donec enim diam vulputate ut pharetra. Sed cras ornare arcu dui.
    
Dui accumsan sit amet nulla. Euismod elementum nisi quis eleifend quam adipiscing vitae proin sagittis. Eleifend mi in nulla posuere sollicitudin aliquam ultrices. Tortor at risus viverra adipiscing at in. Mauris sit amet massa vitae tortor condimentum. Erat pellentesque adipiscing commodo elit. Non enim praesent elementum facilisis. Feugiat vivamus at augue eget arcu dictum varius. Porta non pulvinar neque laoreet. Odio pellentesque diam volutpat commodo sed egestas egestas fringilla. Feugiat nisl pretium fusce id velit ut tortor. Gravida quis blandit turpis cursus in hac habitasse platea dictumst. Id donec ultrices tincidunt arcu non sodales neque sodales ut. At augue eget arcu dictum varius duis. Duis ut diam quam nulla porttitor massa id neque aliquam. Nunc lobortis mattis aliquam faucibus. Tincidunt eget nullam non nisi est. Vestibulum morbi blandit cursus risus at ultrices. Facilisi morbi tempus iaculis urna id volutpat lacus laoreet non.
    
Euismod lacinia at quis risus sed. Enim ut sem viverra aliquet eget sit amet tellus. Eu nisl nunc mi ipsum faucibus vitae. Penatibus et magnis dis parturient montes nascetur ridiculus. Eget gravida cum sociis natoque penatibus et. Feugiat vivamus at augue eget arcu dictum varius duis at. Vitae congue mauris rhoncus aenean. Dictumst quisque sagittis purus sit amet volutpat. Mauris pellentesque pulvinar pellentesque habitant morbi tristique senectus et. Malesuada proin libero nunc consequat. At tempor commodo ullamcorper a lacus. Suspendisse potenti nullam ac tortor.
`;

  initialize() {
    // set dimensions
    this.setWidth(400);
    this.setHeight(400);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Push Error');
      // text list
      this.textList.setWidth(this.getWidth());
      this.textList.setHeight(this.getHeight());
      this.textList.initialize();
  
      this.textList.setX(0);
      this.textList.setY(0);
  
      this.addComponent(this.textList);

      this.setText(this.testText);
  }

  setText(text: string) {
    this.textList.setItemsFromText(text);
  }

  periodicUpdate() {
    //
  }
}
