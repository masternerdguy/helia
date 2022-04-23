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
    // locking during pre-alpha - todo: remove at alpha (this isn't foolproof but it should be sufficient for its purpose)
    const queryString = getQueryStringVars();
    let safe = false;

    if (window.location.href.indexOf('localhost') > 0) {
      safe = true;
    } else {
      for (const k of queryString) {
        if (k == 'safety') {
          if (queryString[k] == '2058') {
            safe = true;
          }
        }
      }
    }

    if (!safe) {
      // locked out
      $('body').remove();
      alert('Not yet in open alpha - coming soon :)');
    } else {
      // normal initialization - keep after removing lock
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
}

function getQueryStringVars(): string[] {
  const pairs: string[] = [];
  const keys = window.location.href
    .slice(window.location.href.indexOf('?') + 1)
    .split('&');

  for (let i = 0; i < keys.length; i++) {
    const key = keys[i].split('=');

    pairs.push(key[0]);
    pairs[key[0]] = key[1];
  }

  return pairs;
}
