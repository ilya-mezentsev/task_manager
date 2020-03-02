import { Component, OnInit } from '@angular/core';
import {ApiRequesterService} from '../../services/api-requester.service';
import {User, UsersListResponse} from '../../interfaces/api-responses';
import {ApiErrorResponse, UserSession} from '../../interfaces/api';
import {ResponseStatus} from '../../interfaces/api';
import {NotifierService} from '../../services/notifier.service';

@Component({
  selector: 'app-group-users-list',
  templateUrl: './group-users-list.component.html',
  styleUrls: ['./group-users-list.component.scss']
})
export class GroupUsersListComponent implements OnInit {
  //public userGroup: UserSession = UserSession.group_id;
  public users: User[] = [];

  constructor(
    private readonly apiRequester: ApiRequesterService,
    private readonly notifierService: NotifierService
  ) {}

  public usersExist(): boolean {
    return this.users.length > 0;
  }

  ngOnInit() {
    this.apiRequester.getUsersList(15)
      .then(usersList => this.processUsersListResponse(usersList))
      .catch(err => {
        console.log(err);
        this.notifierService.send(err);
      });
  }

  private processUsersListResponse(usersList: UsersListResponse | ApiErrorResponse): void {
    if (usersList.status === 'error') {
      console.log(`error while getting users list: ${(usersList as ApiErrorResponse).error_detail}`);
      this.notifierService.send(`error while getting users list: ${(usersList as ApiErrorResponse).error_detail}`);
    } else {
      const users: User[] = (usersList as UsersListResponse).data;
      this.users = users == null
        ? []
        : users;
    }
  }

}
