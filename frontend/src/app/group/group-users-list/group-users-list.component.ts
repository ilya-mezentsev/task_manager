import { Component, OnInit } from '@angular/core';
import {ApiRequesterService} from '../../services/api-requester.service';
import {User, UsersListResponse} from '../../interfaces/api-responses';
import {ApiErrorResponse, UserSession, ResponseStatus} from '../../interfaces/api';
import {NotifierService} from '../../services/notifier.service';
import {StorageService} from '../../services/storage.service';

@Component({
  selector: 'app-group-users-list',
  templateUrl: './group-users-list.component.html',
  styleUrls: ['./group-users-list.component.scss']
})
export class GroupUsersListComponent implements OnInit {
  public user: UserSession = this.storageService.getSession();
  public users: User[] = [];

  constructor(
    private readonly apiRequester: ApiRequesterService,
    private readonly notifierService: NotifierService,
    private readonly storageService: StorageService
  ) {}

  public usersExist(): boolean {
    return this.users.length > 0;
  }

  ngOnInit() {
    this.apiRequester.getUsersList(this.user.id)
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
