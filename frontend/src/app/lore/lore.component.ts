import { Component, OnInit } from '@angular/core';
import { IFactionLoreData } from './factions/iFactionLoreData';
import { WanderersLoreData } from './factions/wanderers';
import * as $ from 'jquery';
import { InterstarLoreData } from './factions/interstar';
import { OriginalsLoreData } from './factions/originals';
import { KingdomLoreData } from './factions/kingdom';
import { AccordLoreData } from './factions/accord';
import { CoalitionLoreData } from './factions/coalition';
import { FederationLoreData } from './factions/federation';
import { AlvacaLoreData } from './factions/alvaca';
import { SanctuaryLoreData } from './factions/sanctuary';
import { BadLoreData } from './factions/bad';
import { FlyLoreData } from './factions/fly';

@Component({
    selector: 'app-lore',
    templateUrl: './lore.component.html',
    styleUrls: ['./lore.component.css'],
    standalone: false
})
export class LoreComponent implements OnInit {
  factionLoreData: IFactionLoreData[] = [
    new OriginalsLoreData(),
    new KingdomLoreData(),
    new AccordLoreData(),
    new CoalitionLoreData(),
    new FederationLoreData(),
    new AlvacaLoreData(),
    new InterstarLoreData(),
    new BadLoreData(),
    new FlyLoreData(),
    new SanctuaryLoreData(),
    new WanderersLoreData(),
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
