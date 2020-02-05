import { Component, OnInit } from '@angular/core';
import {AdminApiRequesterService} from '../../services/admin-api-requester.service';
import {Group, GroupsListResponse} from '../../interfaces/admin-api-responses';
import {ApiErrorResponse} from '../../interfaces/api';
import {ResponseStatus} from '../../interfaces/api';
import {NotifierService} from '../../services/notifier.service';

@Component({
  selector: 'app-all-groups',
  templateUrl: './all-groups.component.html',
  styleUrls: ['./all-groups.component.scss']
})
export class AllGroupsComponent implements OnInit {
  public groups: Group[] = [];

  constructor(
    private readonly adminApiRequester: AdminApiRequesterService,
    private readonly notifierService: NotifierService
  ) {}

  public groupsExist(): boolean {
    return this.groups.length > 0;
  }

  public deleteGroupFromHtml(groupId): void {

  }

  public deleteGroup(groupId: number): void {
    const r = this.adminApiRequester.deleteGroupById(2);
    r.then(result => {
      if (result.status === ResponseStatus.Ok) {
        this.groups = this.groups.filter(group => group.id !== groupId);
      } else {
        this.notifierService.send(`${(result as ApiErrorResponse).error_detail}`);
        return Promise.reject((result as ApiErrorResponse).error_detail);
      }
    });
  }

  ngOnInit() {
    this.adminApiRequester.getGroupsList()
      .then(groupsList => this.processGroupsListResponse(groupsList))
      .catch(err => {
        console.log(err);
        this.notifierService.send(err);
      });
  }

  private processGroupsListResponse(groupsList: GroupsListResponse | ApiErrorResponse): void {
    if (groupsList.status === 'error') {
      console.log(`error while getting groups list: ${(groupsList as ApiErrorResponse).error_detail}`);
      this.notifierService.send(`error while getting groups list: ${(groupsList as ApiErrorResponse).error_detail}`)
    } else {
      const groups: Group[] = (groupsList as GroupsListResponse).data;
      this.groups = groups == null
        ? []
        : groups;
    }
  }
}
