import { Component, OnInit } from '@angular/core';
import {AdminApiRequesterService} from '../../services/admin-api-requester.service';
import {Group, GroupsListResponse} from '../../interfaces/admin-api-responses';
import {ApiErrorResponse} from '../../interfaces/api';

@Component({
  selector: 'app-groups-list',
  templateUrl: './groups-list.component.html',
  styleUrls: ['./groups-list.component.scss']
})
export class GroupsListComponent implements OnInit {
  public groups: Group[] = [];

  constructor(
    private readonly adminAdiRequester: AdminApiRequesterService
  ) {}

  public groupsExist(): boolean {
    return this.groups.length > 0;
  }

  ngOnInit() {
    this.adminAdiRequester.getGroupsList()
      .then(groupsList => this.processGroupsListResponse(groupsList))
      .catch(err => console.log(err));
  }

  private processGroupsListResponse(groupsList: GroupsListResponse | ApiErrorResponse): void {
    if (groupsList.status === 'error') {
      console.log(`error while getting groups list: ${(groupsList as ApiErrorResponse).error_detail}`);
    } else {
      const groups: Group[] = (groupsList as GroupsListResponse).data;
      this.groups = groups == null
        ? []
        : groups;
    }
  }
}
