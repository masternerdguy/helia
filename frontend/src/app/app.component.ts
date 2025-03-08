import { Component } from '@angular/core';
import { AppService } from './app.service';
import * as $ from 'jquery';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  standalone: false,
})
export class AppComponent {
  title = 'Project Helia';

  constructor(private appService: AppService) {}

  public ngOnInit() {
    // capture service reference
    const appService = this.appService;

    $(function () {
      // handle for health check promise
      let pid: any = undefined;

      // local function to check server health
      async function checkHealth(appService: AppService) {
        // get status display area
        const statusArea = $('#server-status-display');

        try {
          // skip if display area is gone
          if (!statusArea[0]) {
            // stop future checks
            clearInterval(pid);

            return;
          }

          // do health check
          const o = await appService.health();

          // update with result
          statusArea.text(o);
        } catch {
          // server is likely down
          statusArea.text('Server offline');
        }
      }

      // enable menu links
      $('.menu-item').on('click', function () {
        const e = $(this);
        const href = e.attr('href');

        window.location.href = href;
      });

      // enable logo link
      $('#header-branding').on('click', function () {
        window.location.href = '/';
      });

      // do initial health check
      checkHealth(appService);

      // do regular health checks
      pid = setInterval(() => {
        checkHealth(appService);
      }, 1000);
    });
  }
}
