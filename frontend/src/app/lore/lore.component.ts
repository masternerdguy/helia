import { Component, OnInit } from '@angular/core';
import { IFactionLoreData } from './factions/iFactionLoreData';
import { WanderersLoreData } from './factions/wanderers';
import * as $ from 'jquery';
import { InterstarLoreData } from './factions/interstar';

@Component({
  selector: 'app-lore',
  templateUrl: './lore.component.html',
  styleUrls: ['./lore.component.css'],
})
export class LoreComponent implements OnInit {
  factionLoreData: IFactionLoreData[] = [
    new InterstarLoreData(),
    new WanderersLoreData()
  ];

  constructor() {}

  ngOnInit(): void {
    $(function () {
      $('.lore-header').on('click', function () {
        const section = $(this).closest('.lore-container');
        section.find('.lore-description ').toggle();
      });
    });
  }
}
