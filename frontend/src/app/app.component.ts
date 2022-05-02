import { Component } from '@angular/core';
import * as $ from 'jquery';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  title = 'Project Helia';

  public ngOnInit() {
    $(function () {
      $('.menu-item').on('click', function () {
        const e = $(this);
        const href = e.attr('href');

        window.location.href = href;
      });

      $('#header-branding').on('click', function () {
        window.location.href = '/';
      });
    });
  }
}
