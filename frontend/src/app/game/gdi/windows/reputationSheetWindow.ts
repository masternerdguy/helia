import { GDIWindow } from '../base/gdiWindow';
import { FontSize, GDIStyle } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import {
  GetFactionCache,
  GetFactionCacheEntry,
  GetPlayerFactionRelationshipCache,
} from '../../wsModels/shared';
import { Faction } from '../../engineModels/faction';
import { WSPlayerFactionRelationship } from '../../wsModels/entities/wsPlayerFaction';
import { GDITabPane } from '../components/gdiTabPane';
import { GDILabel } from '../components/gdiLabel';
import { Player } from '../../engineModels/player';
import { WsService } from '../../ws.service';
import { GDIButton } from '../components/gdiButton';
import { GDIOverlay } from '../components/gdiOverlay';
import { GDIInput } from '../components/gdiInput';
import { ClientNewFaction } from '../../wsModels/bodies/newFaction';
import { MessageTypes } from '../../wsModels/gameMessage';
import { ClientLeaveFaction } from '../../wsModels/bodies/leaveFaction';
import { ClientApplyToFaction } from '../../wsModels/bodies/applyToFaction';
import { ServerApplicationsUpdate } from '../../wsModels/bodies/applicationsUpdate';
import { ViewApplications as ClientViewApplications } from '../../wsModels/bodies/viewApplications';
import { WSApplication } from '../../wsModels/entities/wsApplication';
import { ClientApproveApplication } from '../../wsModels/bodies/approveApplication';
import { ClientRejectApplication } from '../../wsModels/bodies/rejectApplication';
import { ClientKickMember } from '../../wsModels/bodies/kickMember';
import { ServerMembersUpdate } from '../../wsModels/bodies/membersUpdate';
import { WSMember } from '../../wsModels/entities/wsMember';
import { ClientViewMembers } from '../../wsModels/bodies/viewMembers';

export class ReputationSheetWindow extends GDIWindow {
  private player: Player;
  private lastFactionId: string;
  private isFactionOwner: boolean;
  private lastApplicationView: number = 0;
  private lastMemberView: number = 0;
  private wsSvc: WsService;

  // modal base
  private modalOverlay: GDIOverlay = new GDIOverlay();
  private modalInput: GDIInput = new GDIInput();

  // tab pane
  private tabs = new GDITabPane();

  // "reputation" tab
  private factionList = new GDIList();
  private actionList = new GDIList();
  private infoList = new GDIList();

  // "my faction" tab (in NPC faction)
  private npcFactionLabel = new GDILabel();
  private createFactionButton = new GDIButton();

  // "my faction" tab (in player faction)
  private playerFactionLabel = new GDILabel();
  private leaveFactionButton = new GDIButton();
  private viewApplicantsButton = new GDIButton();
  private viewMembersButton = new GDIButton();

  // "new faction" modal form
  private newFactionNameLabel = new GDILabel();
  private newFactionNameInput = new GDIInput();
  private newFactionTickerLabel = new GDILabel();
  private newFactionTickerInput = new GDIInput();
  private newFactionDescriptionLabel = new GDILabel();
  private newFactionDescriptionInput = new GDIInput();
  private newFactionSubmitButton = new GDIButton();
  private newFactionCancelButton = new GDIButton();

  // "view applicants" modal form
  private applicantList = new GDIList();
  private applicantActionList = new GDIList();
  private applicantInfoList = new GDIList();
  private viewApplicantCloseButton = new GDIButton();

  // "view members" modal form
  private memberList = new GDIList();
  private memberActionList = new GDIList();
  private memberInfoList = new GDIList();
  private viewMembersCloseButton = new GDIButton();

  initialize() {
    // set dimensions
    this.setWidth(400);
    this.setHeight(420);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Reputation Sheet');

    this.setOnShow(() => {
      // defer refresh
      setTimeout(() => (this.lastFactionId = ''), 10);
    });

    // tab list
    this.tabs.setWidth(this.getWidth());
    this.tabs.setHeight(this.getHeight());
    this.tabs.initialize();

    this.tabs.setX(0);
    this.tabs.setY(0);
    this.tabs.setSelectedTab('Reputation');

    this.addComponent(this.tabs);

    // modal base and generic input
    this.modalOverlay.setWidth(this.getWidth());
    this.modalOverlay.setHeight(this.getHeight());
    this.modalOverlay.setX(0);
    this.modalOverlay.setY(0);
    this.modalOverlay.initialize();

    const fontSize = GDIStyle.getUnderlyingFontSize(FontSize.large);
    this.modalInput.setWidth(100);
    this.modalInput.setHeight(Math.round(fontSize + 0.5));
    this.modalInput.setX(this.getWidth() / 2 - this.modalInput.getWidth() / 2);
    this.modalInput.setY(
      this.getHeight() / 2 - this.modalInput.getHeight() / 2
    );

    this.modalInput.setFont(FontSize.large);
    this.modalInput.initialize();

    // pack tabs
    this.packReputationTab();
    this.packMyFactionTab();

    // pack modal views
    this.packViewApplicantsModal();
    this.packViewMembersModal();

    // pack modal forms
    this.packNewFactionModalForm();
  }

  private packViewMembersModal() {
    // member list
    this.memberList.setWidth(300);
    this.memberList.setHeight(200 + GDIStyle.tabHandleHeight);
    this.memberList.initialize();

    this.memberList.setX(0);
    this.memberList.setY(0);

    this.memberList.setFont(FontSize.normal);
    this.memberList.setOnClick((item) => {
      // get row
      const o = item as MemberInfoRow;

      if (o.member != null) {
        // update actions
        const actions: any[] = [];

        for (const a of o.actions) {
          actions.push({
            listString: () => `${a}`,
          });
        }

        this.memberActionList.setItems(actions);

        // update detailed info
        const details = this.buildMemberDetails(o);
        this.memberInfoList.setItems(details);
      } else {
        this.memberActionList.setItems([]);
        this.memberInfoList.setItems([]);
      }
    });

    // action list
    this.memberActionList.setWidth(100);
    this.memberActionList.setHeight(180 + GDIStyle.tabHandleHeight);
    this.memberActionList.initialize();

    this.memberActionList.setX(300);
    this.memberActionList.setY(0);

    this.memberActionList.setFont(FontSize.normal);
    this.memberActionList.setOnClick((item) => {
      const action = item.listString();
      const member = this.memberList.getSelectedItem() as MemberInfoRow;

      if (action == 'Kick') {
        // send request to kick member
        const b = new ClientKickMember();

        b.sid = this.wsSvc.sid;
        b.userId = member.member.id;

        this.wsSvc.sendMessage(MessageTypes.KickMember, b);

        // request refresh
        const now = Date.now();
        setTimeout(() => this.refreshMembers(now), 0);

        // reset lists
        this.memberActionList.setItems([]);
        this.memberList.setItems([]);
        this.memberInfoList.setItems([]);
      }
    });

    // info list
    this.memberInfoList.setWidth(400);
    this.memberInfoList.setHeight(190 - GDIStyle.tabHandleHeight);
    this.memberInfoList.initialize();

    this.memberInfoList.setX(0);
    this.memberInfoList.setY(200);

    this.memberInfoList.setFont(FontSize.normal);
    this.memberInfoList.setOnClick(() => {});

    // close button
    this.viewMembersCloseButton.setWidth(this.getWidth() * 0.5);
    this.viewMembersCloseButton.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );
    this.viewMembersCloseButton.initialize();

    this.viewMembersCloseButton.setText('Close');
    this.viewMembersCloseButton.setFont(FontSize.normal);

    this.viewMembersCloseButton.setX(this.getWidth() * 0.25);
    this.viewMembersCloseButton.setY(this.getHeight() - 40);

    this.viewMembersCloseButton.setOnClick(() => {
      // close view members modal
      this.hideViewMembersModal();
    });
  }

  private packViewApplicantsModal() {
    // applicant list
    this.applicantList.setWidth(300);
    this.applicantList.setHeight(200 + GDIStyle.tabHandleHeight);
    this.applicantList.initialize();

    this.applicantList.setX(0);
    this.applicantList.setY(0);

    this.applicantList.setFont(FontSize.normal);
    this.applicantList.setOnClick((item) => {
      // get row
      const o = item as ApplicantInfoRow;

      if (o.applicant != null) {
        // update actions
        const actions: any[] = [];

        for (const a of o.actions) {
          actions.push({
            listString: () => `${a}`,
          });
        }

        this.applicantActionList.setItems(actions);

        // update detailed info
        const details = this.buildApplicantDetails(o);
        this.applicantInfoList.setItems(details);
      } else {
        this.applicantActionList.setItems([]);
        this.applicantInfoList.setItems([]);
      }
    });

    // action list
    this.applicantActionList.setWidth(100);
    this.applicantActionList.setHeight(180 + GDIStyle.tabHandleHeight);
    this.applicantActionList.initialize();

    this.applicantActionList.setX(300);
    this.applicantActionList.setY(0);

    this.applicantActionList.setFont(FontSize.normal);
    this.applicantActionList.setOnClick((item) => {
      const action = item.listString();
      const application =
        this.applicantList.getSelectedItem() as ApplicantInfoRow;

      if (action == 'Approve') {
        // send request to approve application
        const b = new ClientApproveApplication();

        b.sid = this.wsSvc.sid;
        b.userId = application.applicant.id;

        this.wsSvc.sendMessage(MessageTypes.ApproveApplication, b);

        // request refresh
        const now = Date.now();
        setTimeout(() => this.refreshApplicants(now), 0);

        // reset lists
        this.applicantActionList.setItems([]);
        this.applicantList.setItems([]);
        this.applicantInfoList.setItems([]);
      } else if (action == 'Reject') {
        // send request to reject application
        const b = new ClientRejectApplication();

        b.sid = this.wsSvc.sid;
        b.userId = application.applicant.id;

        this.wsSvc.sendMessage(MessageTypes.RejectApplication, b);

        // request refresh
        const now = Date.now();
        setTimeout(() => this.refreshApplicants(now), 0);

        // reset lists
        this.applicantActionList.setItems([]);
        this.applicantList.setItems([]);
        this.applicantInfoList.setItems([]);
      }
    });

    // info list
    this.applicantInfoList.setWidth(400);
    this.applicantInfoList.setHeight(190 - GDIStyle.tabHandleHeight);
    this.applicantInfoList.initialize();

    this.applicantInfoList.setX(0);
    this.applicantInfoList.setY(200);

    this.applicantInfoList.setFont(FontSize.normal);
    this.applicantInfoList.setOnClick(() => {});

    // close button
    this.viewApplicantCloseButton.setWidth(this.getWidth() * 0.5);
    this.viewApplicantCloseButton.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );
    this.viewApplicantCloseButton.initialize();

    this.viewApplicantCloseButton.setText('Close');
    this.viewApplicantCloseButton.setFont(FontSize.normal);

    this.viewApplicantCloseButton.setX(this.getWidth() * 0.25);
    this.viewApplicantCloseButton.setY(this.getHeight() - 40);

    this.viewApplicantCloseButton.setOnClick(() => {
      // close view applicants modal
      this.hideViewApplicantsModal();
    });
  }

  private packNewFactionModalForm() {
    const inputFontSize = GDIStyle.getUnderlyingFontSize(FontSize.large);

    // new faction name label
    this.newFactionNameLabel.setWidth(this.getWidth());
    this.newFactionNameLabel.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );
    this.newFactionNameLabel.initialize();

    this.newFactionNameLabel.setText('Faction Name');
    this.newFactionNameLabel.setFont(FontSize.normal);

    this.newFactionNameLabel.setX(0);
    this.newFactionNameLabel.setY(GDIStyle.tabHandleHeight);

    // new faction name input
    this.newFactionNameInput.setWidth(this.getWidth() * 0.95);
    this.newFactionNameInput.setHeight(inputFontSize);

    this.newFactionNameInput.initialize();
    this.newFactionNameInput.setFont(FontSize.large);

    this.newFactionNameInput.setX(this.getWidth() * 0.025);
    this.newFactionNameInput.setY(
      this.newFactionNameLabel.getY() +
        this.newFactionNameLabel.getHeight() +
        10
    );

    // new faction description label
    this.newFactionDescriptionLabel.setWidth(this.getWidth());
    this.newFactionDescriptionLabel.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );
    this.newFactionDescriptionLabel.initialize();

    this.newFactionDescriptionLabel.setText('Brief Description');
    this.newFactionDescriptionLabel.setFont(FontSize.normal);

    this.newFactionDescriptionLabel.setX(0);
    this.newFactionDescriptionLabel.setY(
      this.newFactionNameInput.getY() +
        this.newFactionNameInput.getHeight() +
        10
    );

    // new faction description input
    this.newFactionDescriptionInput.setWidth(this.getWidth() * 0.95);
    this.newFactionDescriptionInput.setHeight(inputFontSize);

    this.newFactionDescriptionInput.initialize();
    this.newFactionDescriptionInput.setFont(FontSize.large);

    this.newFactionDescriptionInput.setX(this.getWidth() * 0.025);
    this.newFactionDescriptionInput.setY(
      this.newFactionDescriptionLabel.getY() +
        this.newFactionDescriptionLabel.getHeight() +
        10
    );

    // new faction ticker label
    this.newFactionTickerLabel.setWidth(this.getWidth());
    this.newFactionTickerLabel.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );
    this.newFactionTickerLabel.initialize();

    this.newFactionTickerLabel.setText('Ticker');
    this.newFactionTickerLabel.setFont(FontSize.normal);

    this.newFactionTickerLabel.setX(0);
    this.newFactionTickerLabel.setY(
      this.newFactionDescriptionInput.getY() +
        this.newFactionDescriptionInput.getHeight() +
        10
    );

    // new faction ticker input
    this.newFactionTickerInput.setWidth(this.getWidth() * 0.95);
    this.newFactionTickerInput.setHeight(inputFontSize);

    this.newFactionTickerInput.initialize();
    this.newFactionTickerInput.setFont(FontSize.large);

    this.newFactionTickerInput.setX(this.getWidth() * 0.025);
    this.newFactionTickerInput.setY(
      this.newFactionTickerLabel.getY() +
        this.newFactionTickerLabel.getHeight() +
        10
    );

    // submit button
    this.newFactionSubmitButton.setWidth(this.getWidth() * 0.5);
    this.newFactionSubmitButton.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );
    this.newFactionSubmitButton.initialize();

    this.newFactionSubmitButton.setText('Submit');
    this.newFactionSubmitButton.setFont(FontSize.normal);

    this.newFactionSubmitButton.setX(this.getWidth() * 0.25);
    this.newFactionSubmitButton.setY(this.getHeight() - 130);

    this.newFactionSubmitButton.setOnClick(() => {
      // send request to create new faction
      const b = new ClientNewFaction();

      b.sid = this.wsSvc.sid;
      b.name = this.newFactionNameInput.getText();
      b.description = this.newFactionDescriptionInput.getText();
      b.ticker = this.newFactionTickerInput.getText();

      this.wsSvc.sendMessage(MessageTypes.CreateNewFaction, b);

      // close form
      this.hideNewFactionFormModal();

      // close window
      this.setHidden(true);
    });

    // cancel button
    this.newFactionCancelButton.setWidth(this.getWidth() * 0.5);
    this.newFactionCancelButton.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );
    this.newFactionCancelButton.initialize();

    this.newFactionCancelButton.setText('Cancel');
    this.newFactionCancelButton.setFont(FontSize.normal);

    this.newFactionCancelButton.setX(this.getWidth() * 0.25);
    this.newFactionCancelButton.setY(this.getHeight() - 100);

    this.newFactionCancelButton.setOnClick(() => {
      // close new faction form
      this.hideNewFactionFormModal();
    });
  }

  private packMyFactionTab() {
    // npc faction indicator label
    this.npcFactionLabel.setWidth(this.getWidth());
    this.npcFactionLabel.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );
    this.npcFactionLabel.initialize();

    this.npcFactionLabel.setText(
      'You are currently a member of an NPC faction.'
    );
    this.npcFactionLabel.setFont(FontSize.normal);

    this.npcFactionLabel.setX(0);
    this.npcFactionLabel.setY(GDIStyle.tabHandleHeight);

    // player faction indicator label
    this.playerFactionLabel.setWidth(this.getWidth());
    this.playerFactionLabel.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );
    this.playerFactionLabel.initialize();

    this.playerFactionLabel.setText(
      'You are currently a member of a player faction.'
    );
    this.playerFactionLabel.setFont(FontSize.normal);

    this.playerFactionLabel.setX(0);
    this.playerFactionLabel.setY(GDIStyle.tabHandleHeight);

    // create faction button
    this.createFactionButton.setWidth(this.getWidth() * 0.5);
    this.createFactionButton.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );
    this.createFactionButton.initialize();

    this.createFactionButton.setText('Create Faction');
    this.createFactionButton.setFont(FontSize.normal);

    this.createFactionButton.setX(this.getWidth() * 0.25);
    this.createFactionButton.setY(100 + GDIStyle.tabHandleHeight);

    this.createFactionButton.setOnClick(() => {
      // show new faction form
      this.showNewFactionFormModal();

      // reset form inputs
      this.resetNewFactionFormInputs();
    });

    // leave faction button
    this.leaveFactionButton.setWidth(this.getWidth() * 0.5);
    this.leaveFactionButton.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );

    this.leaveFactionButton.initialize();

    this.leaveFactionButton.setText('Leave Faction');
    this.leaveFactionButton.setFont(FontSize.normal);

    this.leaveFactionButton.setX(this.getWidth() * 0.25);
    this.leaveFactionButton.setY(100 + GDIStyle.tabHandleHeight);

    this.leaveFactionButton.setBorderColor('red');
    this.leaveFactionButton.setTextColor('red');

    this.leaveFactionButton.setOnClick(() => {
      // send request to leave faction
      const b = new ClientLeaveFaction();

      b.sid = this.wsSvc.sid;

      this.wsSvc.sendMessage(MessageTypes.LeaveFaction, b);
    });

    // view applicants button
    this.viewApplicantsButton.setWidth(this.getWidth() * 0.5);
    this.viewApplicantsButton.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );

    this.viewApplicantsButton.initialize();

    this.viewApplicantsButton.setText('View Applicants');
    this.viewApplicantsButton.setFont(FontSize.normal);

    this.viewApplicantsButton.setX(this.getWidth() * 0.25);
    this.viewApplicantsButton.setY(
      this.leaveFactionButton.getY() + this.leaveFactionButton.getHeight() + 10
    );

    this.viewApplicantsButton.setOnClick(() => {
      // show view applicants modal
      this.showViewApplicantsModal();
    });

    // view members button
    this.viewMembersButton.setWidth(this.getWidth() * 0.5);
    this.viewMembersButton.setHeight(
      Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2)
    );

    this.viewMembersButton.initialize();

    this.viewMembersButton.setText('View Members');
    this.viewMembersButton.setFont(FontSize.normal);

    this.viewMembersButton.setX(this.getWidth() * 0.25);
    this.viewMembersButton.setY(
      this.viewApplicantsButton.getY() +
        this.viewApplicantsButton.getHeight() +
        10
    );

    this.viewMembersButton.setOnClick(() => {
      // show view applicants modal
      this.showViewMembersModal();
    });
  }

  private packReputationTab() {
    // faction list
    this.factionList.setWidth(300);
    this.factionList.setHeight(200 + GDIStyle.tabHandleHeight);
    this.factionList.initialize();

    this.factionList.setX(0);
    this.factionList.setY(0);

    this.factionList.setFont(FontSize.normal);
    this.factionList.setOnClick((item) => {
      // get faction row
      const o = item as FactionRepViewRow;

      if (o.faction != null) {
        // update actions
        this.actionList.setItems(o.actions);

        // update detailed info
        const details = this.buildDetails(o);
        this.infoList.setItems(details);
      } else {
        this.actionList.setItems([]);
        this.infoList.setItems([]);
      }
    });

    this.tabs.addComponent(this.factionList, 'Reputation');

    // action list
    this.actionList.setWidth(100);
    this.actionList.setHeight(200 + GDIStyle.tabHandleHeight);
    this.actionList.initialize();

    this.actionList.setX(300);
    this.actionList.setY(0);

    this.actionList.setFont(FontSize.normal);
    this.actionList.setOnClick((item) => {
      const action = item.listString();

      if (action == 'Apply') {
        // send request to apply to join faction
        const b = new ClientApplyToFaction();

        b.sid = this.wsSvc.sid;
        b.factionId = item.faction.id;

        this.wsSvc.sendMessage(MessageTypes.ApplyToFaction, b);
      }
    });

    this.tabs.addComponent(this.actionList, 'Reputation');

    // info list
    this.infoList.setWidth(400);
    this.infoList.setHeight(220 - GDIStyle.tabHandleHeight);
    this.infoList.initialize();

    this.infoList.setX(0);
    this.infoList.setY(200);

    this.infoList.setFont(FontSize.normal);
    this.infoList.setOnClick(() => {});

    this.tabs.addComponent(this.infoList, 'Reputation');
  }

  private buildDetails(r: FactionRepViewRow): FactionInfoViewRow[] {
    const relationships = r.faction.relationships.sort(
      (a, b) => a.standingValue - b.standingValue
    );

    const factions = GetFactionCache().sort((a, b) => {
      // get standing entries
      const aStanding = relationships.filter((x) => x.factionId == a.id);
      const bStanding = relationships.filter((x) => x.factionId == b.id);

      // extract values
      var aVal = 0;
      var bVal = 0;

      if (aStanding.length > 0) {
        aVal = aStanding[0].standingValue;
      }

      if (bStanding.length > 0) {
        bVal = bStanding[0].standingValue;
      }

      // sort by standing desc
      return bVal - aVal;
    });

    const rows: string[] = [];

    // basic info
    rows.push('Basic');
    rows.push(infoKeyValueString('Name', r.faction.name));
    rows.push(infoKeyValueString('Ticker', `[${r.faction.ticker}]`));
    rows.push('');

    // relationships (NPC)
    if (r.faction.isNPC) {
      rows.push('Liked By');
      for (const f of factions) {
        // only NPC factions
        if (!f.isNPC) {
          continue;
        }

        if (f.id != r.faction.id) {
          // find relationship
          for (const rel of f.relationships) {
            if (rel.factionId != r.faction.id) {
              continue;
            }

            if (rel.standingValue > 0) {
              let openHostileFlag = '';

              if (rel.openlyHostile) {
                openHostileFlag = '⚔';
              }

              rows.push(
                infoKeyValueString(
                  f.name,
                  `[${rel.standingValue.toFixed(2)}] ` + openHostileFlag
                )
              );
            }
          }
        }
      }

      rows.push('');

      rows.push('Disliked By');
      for (const f of factions) {
        if (f.id != r.faction.id) {
          // only NPC factions
          if (!f.isNPC) {
            continue;
          }

          // find relationship
          for (const rel of f.relationships) {
            if (rel.factionId != r.faction.id) {
              continue;
            }

            if (rel.standingValue < 0) {
              let openHostileFlag = '';

              if (rel.openlyHostile) {
                openHostileFlag = '⚔';
              }

              rows.push(
                infoKeyValueString(
                  f.name,
                  `[${rel.standingValue.toFixed(2)}] ` + openHostileFlag
                )
              );
            }
          }
        }
      }
    } else {
      // relationships (player)
      rows.push('Liked By');
      for (const f of r.faction.relationships) {
        if (f.factionId != r.faction.id) {
          if (f.standingValue > 0) {
            let openHostileFlag = '';

            if (f.openlyHostile) {
              openHostileFlag = '⚔';
            }

            const fc = GetFactionCacheEntry(f.factionId);

            rows.push(
              infoKeyValueString(
                fc.name,
                `[${f.standingValue.toFixed(2)}] ` + openHostileFlag
              )
            );
          }
        }
      }

      rows.push('');

      rows.push('Disliked By');
      for (const f of r.faction.relationships) {
        if (f.factionId != r.faction.id) {
          if (f.standingValue < 0) {
            let openHostileFlag = '';

            if (f.openlyHostile) {
              openHostileFlag = '⚔';
            }

            const fc = GetFactionCacheEntry(f.factionId);

            rows.push(
              infoKeyValueString(
                fc.name,
                `[${f.standingValue.toFixed(2)}] ` + openHostileFlag
              )
            );
          }
        }
      }
    }

    // description
    rows.push('');

    if (r.faction.isNPC) {
      rows.push('Encyclopedia Summary');
    } else {
      rows.push('Player-provided Description');
    }

    rows.push('');

    const displayDescription = this.infoList.breakText(
      r.faction.description ?? ''
    );

    for (const i of displayDescription) {
      rows.push(i.text);
    }

    // convert to display rows
    const dRows: FactionInfoViewRow[] = [];

    for (const r of rows) {
      dRows.push(this.infoRowFromString(r));
    }

    // return converted rows
    return dRows;
  }

  private infoRowFromString(s: string): FactionInfoViewRow {
    return {
      listString: () => s,
    };
  }

  private applicantInfoViewRowFromString(s: string): ApplicantInfoViewRow {
    return {
      listString: () => s,
    };
  }

  private memberInfoViewRowFromString(s: string): MemberInfoViewRow {
    return {
      listString: () => s,
    };
  }

  private buildApplicantDetails(r: ApplicantInfoRow): ApplicantInfoViewRow[] {
    const rows: string[] = [];

    // basic info
    rows.push('Basic');
    rows.push(infoKeyValueString('Name', r.applicant.name));
    rows.push(infoKeyValueString('Faction', `${r.applicant.faction}`));
    rows.push(infoKeyValueString('Ticker', `[${r.applicant.ticker}]`));
    rows.push('');

    // convert to display rows
    const dRows: ApplicantInfoViewRow[] = [];

    for (const r of rows) {
      dRows.push(this.applicantInfoViewRowFromString(r));
    }

    // return converted rows
    return dRows;
  }

  private buildMemberDetails(r: MemberInfoRow): MemberInfoViewRow[] {
    const rows: string[] = [];

    // basic info
    rows.push('Basic');
    rows.push(infoKeyValueString('Name', r.member.name));
    rows.push('');

    // convert to display rows
    const dRows: MemberInfoViewRow[] = [];

    for (const r of rows) {
      dRows.push(this.memberInfoViewRowFromString(r));
    }

    // return converted rows
    return dRows;
  }

  setWsService(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  setPlayer(player: Player) {
    this.player = player;
    const faction = player.getFaction();

    // check for faction change (or initial presentation)
    if (this.lastFactionId != faction.id) {
      this.lastFactionId = faction.id;

      if (faction.isNPC) {
        this.showNPCComponentsOnMyFactionTab();
        this.hidePlayerComponentsOnMyFactionTab();
      } else {
        this.hideNPCComponentsOnMyFactionTab();
        this.showPlayerComponentsOnMyFactionTab();
      }
    }
  }

  periodicUpdate() {
    if (!this.isHidden()) {
      // get player-faction relationships
      const playerFactionRelationships = GetPlayerFactionRelationshipCache();

      // get factions
      const factions = GetFactionCache().sort((a, b) => {
        // get standing entries
        const aStanding = playerFactionRelationships.filter(
          (x) => x.factionId == a.id
        );

        const bStanding = playerFactionRelationships.filter(
          (x) => x.factionId == b.id
        );

        // extract values
        var aVal = 0;
        var bVal = 0;

        if (aStanding.length > 0) {
          aVal = aStanding[0].standingValue;
        }

        if (bStanding.length > 0) {
          bVal = bStanding[0].standingValue;
        }

        // sort by standing desc
        return bVal - aVal;
      });

      // build rows (NPC factions)
      const factionRows: FactionRepViewRow[] = [];

      for (const f of factions.filter(
        (f) => f.isNPC && f.id != '42b937ad-0000-46e9-9af9-fc7dbf878e6a'
      )) {
        let playerRel: WSPlayerFactionRelationship = null;

        // find relationship to player
        for (const rel of playerFactionRelationships) {
          if (rel.factionId == f.id) {
            playerRel = rel;
            break;
          }
        }

        factionRows.push({
          faction: f,
          actions: [],
          listString: () => factionListRowString(playerRel, f),
        });
      }

      // spacer
      factionRows.push({
        faction: null,
        actions: [],
        listString: () => '',
      });

      factionRows.push({
        faction: null,
        actions: [],
        listString: () => 'Player Factions',
      });

      factionRows.push({
        faction: null,
        actions: [],
        listString: () => '===============',
      });

      factionRows.push({
        faction: null,
        actions: [],
        listString: () => '',
      });

      // build rows (player factions)
      for (const f of factions.filter(
        (f) => !f.isNPC && f.id != '42b937ad-0000-46e9-9af9-fc7dbf878e6a'
      )) {
        let playerRel: WSPlayerFactionRelationship = null;

        // find relationship to player
        for (const rel of playerFactionRelationships) {
          if (rel.factionId == f.id) {
            playerRel = rel;
            break;
          }
        }

        if (!playerRel) {
          // stub a zero standing entry for unknown faction
          playerRel = {
            factionId: f.id,
            openlyHostile: false,
            standingValue: 0,
            isMember: f.id == this.player.getFaction().id,
          };
        }

        const actionList = [];

        if (f.isJoinable && !f.isNPC) {
          actionList.push({
            listString: () => 'Apply',
            faction: f,
          });
        }

        factionRows.push({
          faction: f,
          actions: actionList,
          listString: () => factionListRowString(playerRel, f),
        });
      }

      // update faction list
      const sp = this.factionList.getScroll();
      const si = this.factionList.getSelectedIndex();

      this.factionList.setItems(factionRows);
      this.factionList.setScroll(sp);
      this.factionList.setSelectedIndex(si);

      // check if player is faction owner
      if (this.player.uid == this.player.getFaction().ownerId) {
        // set flag
        this.isFactionOwner = true;

        // check if applications view is stale
        const now = Date.now();

        if (now - this.lastApplicationView > 5000) {
          // request current applicants
          this.refreshApplicants(now);
        }

        // check if members view is stale
        if (now - this.lastMemberView > 5000) {
          // request current members
          this.refreshMembers(now);
        }
      } else {
        this.isFactionOwner = false;
      }
    }
  }

  private refreshApplicants(now: number) {
    const b = new ClientViewApplications();
    b.sid = this.wsSvc.sid;

    this.wsSvc.sendMessage(MessageTypes.ViewApplications, b);
    this.lastApplicationView = now;
  }

  private refreshMembers(now: number) {
    const b = new ClientViewMembers();
    b.sid = this.wsSvc.sid;

    this.wsSvc.sendMessage(MessageTypes.ViewMembers, b);
    this.lastMemberView = now;
  }

  private showNPCComponentsOnMyFactionTab() {
    this.tabs.addComponent(this.npcFactionLabel, 'My Faction');
    this.tabs.addComponent(this.createFactionButton, 'My Faction');
  }

  private hideNPCComponentsOnMyFactionTab() {
    this.tabs.removeComponent(this.npcFactionLabel, 'My Faction');
    this.tabs.removeComponent(this.createFactionButton, 'My Faction');
  }

  private showPlayerComponentsOnMyFactionTab() {
    this.tabs.addComponent(this.playerFactionLabel, 'My Faction');
    this.tabs.addComponent(this.leaveFactionButton, 'My Faction');

    if (this.isFactionOwner) {
      this.tabs.addComponent(this.viewApplicantsButton, 'My Faction');
      this.tabs.addComponent(this.viewMembersButton, 'My Faction');
    }
  }

  private hidePlayerComponentsOnMyFactionTab() {
    this.tabs.removeComponent(this.playerFactionLabel, 'My Faction');
    this.tabs.removeComponent(this.leaveFactionButton, 'My Faction');
    this.tabs.removeComponent(this.viewApplicantsButton, 'My Faction');
    this.tabs.removeComponent(this.viewMembersButton, 'My Faction');
  }

  private showNewFactionFormModal() {
    this.showModalBase();

    this.addComponent(this.newFactionNameLabel);
    this.addComponent(this.newFactionNameInput);
    this.addComponent(this.newFactionDescriptionLabel);
    this.addComponent(this.newFactionDescriptionInput);
    this.addComponent(this.newFactionTickerLabel);
    this.addComponent(this.newFactionTickerInput);
    this.addComponent(this.newFactionSubmitButton);
    this.addComponent(this.newFactionCancelButton);
  }

  private hideNewFactionFormModal() {
    this.hideModalBase();

    this.removeComponent(this.newFactionNameLabel);
    this.removeComponent(this.newFactionNameInput);
    this.removeComponent(this.newFactionDescriptionLabel);
    this.removeComponent(this.newFactionDescriptionInput);
    this.removeComponent(this.newFactionTickerLabel);
    this.removeComponent(this.newFactionTickerInput);
    this.removeComponent(this.newFactionSubmitButton);
    this.removeComponent(this.newFactionCancelButton);
  }

  private showViewApplicantsModal() {
    this.showModalBase();

    this.addComponent(this.applicantList);
    this.addComponent(this.applicantInfoList);
    this.addComponent(this.applicantActionList);
    this.addComponent(this.viewApplicantCloseButton);
  }

  private hideViewApplicantsModal() {
    this.hideModalBase();

    this.removeComponent(this.applicantList);
    this.removeComponent(this.applicantInfoList);
    this.removeComponent(this.applicantActionList);
    this.removeComponent(this.viewApplicantCloseButton);
  }

  private showViewMembersModal() {
    this.showModalBase();

    this.addComponent(this.memberList);
    this.addComponent(this.memberInfoList);
    this.addComponent(this.memberActionList);
    this.addComponent(this.viewMembersCloseButton);
  }

  private hideViewMembersModal() {
    this.hideModalBase();

    this.removeComponent(this.memberList);
    this.removeComponent(this.memberInfoList);
    this.removeComponent(this.memberActionList);
    this.removeComponent(this.viewMembersCloseButton);
  }

  private showModalBase() {
    this.removeComponent(this.tabs);
    this.addComponent(this.modalOverlay);
  }

  private hideModalBase() {
    this.addComponent(this.tabs);
    this.removeComponent(this.modalOverlay);
  }

  private resetNewFactionFormInputs() {
    this.newFactionNameInput.setText('');
    this.newFactionDescriptionInput.setText('');
    this.newFactionTickerInput.setText('');
  }

  syncApplications(msg: ServerApplicationsUpdate) {
    // stash applicant list state
    const scroll = this.applicantList.getScroll();
    const selection = this.applicantList.getSelectedIndex();

    // clear applicant list
    this.applicantList.setItems([]);

    // build new rows
    const rows: ApplicantInfoRow[] = [];

    for (const r of msg.applications) {
      rows.push({
        actions: ['Approve', '', 'Reject'],
        applicant: r,
        listString: () => `${r.name}`,
      });
    }

    // store new rows
    this.applicantList.setItems(rows);

    // restore state
    this.applicantList.setSelectedIndex(selection);
    this.applicantList.setScroll(scroll);
  }

  syncMembers(msg: ServerMembersUpdate) {
    // stash applicant list state
    const scroll = this.memberList.getScroll();
    const selection = this.memberList.getSelectedIndex();

    // clear applicant list
    this.memberList.setItems([]);

    // build new rows
    const rows: MemberInfoRow[] = [];

    for (const r of msg.members) {
      rows.push({
        actions: ['Kick', ''],
        member: r,
        listString: () => `${r.name}`,
      });
    }

    // store new rows
    this.memberList.setItems(rows);

    // restore state
    this.memberList.setSelectedIndex(selection);
    this.memberList.setScroll(scroll);
  }
}

class FactionRepViewRow {
  faction: Faction;
  actions: string[];

  listString: () => string;
}

class FactionInfoViewRow {
  listString: () => string;
}

class ApplicantInfoRow {
  applicant: WSApplication;
  actions: string[];

  listString: () => string;
}

class ApplicantInfoViewRow {
  listString: () => string;
}

class MemberInfoRow {
  member: WSMember;
  actions: string[];

  listString: () => string;
}

class MemberInfoViewRow {
  listString: () => string;
}

function factionListRowString(
  rel: WSPlayerFactionRelationship,
  faction: Faction
): string {
  if (rel == null || faction == null) {
    return;
  }

  // determine whether or not to show hostility flag
  let openHostileFlag = '';

  if (rel.openlyHostile) {
    openHostileFlag = '⚔';
  }

  // determine whether or not the player is a member
  let memberFlag = '';

  if (rel.isMember) {
    memberFlag = '✪';
  }

  // build string
  return `${fixedString(faction.name, 24)} ${fixedString(
    `[${rel.standingValue.toFixed(2)}] `,
    10
  )} ${fixedString(memberFlag, 1)}${fixedString(openHostileFlag, 1)}`;
}

function infoKeyValueString(key: string, value: string) {
  // build string
  return `${fixedString('', 1)} ${fixedString(key, 24)} ${fixedString(
    value,
    24
  )}`;
}

function fixedString(str: string, width: number): string {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str.substr(0, width).padEnd(width);
}
