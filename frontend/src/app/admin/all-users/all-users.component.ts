import { Component, OnInit } from '@angular/core';
import {AdminApiRequesterService} from '../../services/admin-api-requester.service';
import {User, UsersListResponse} from '../../interfaces/admin-api-responses';
import {ApiErrorResponse, ResponseStatus} from '../../interfaces/api';
import {NotifierService} from '../../services/notifier.service';

@Component({
  selector: 'app-all-users',
  templateUrl: './all-users.component.html',
  styleUrls: ['./all-users.component.scss']
})
export class AllUsersComponent implements OnInit {

  public users: User[] = [];

  constructor(
    private readonly adminApiRequester: AdminApiRequesterService,
    private readonly notifierService: NotifierService
  ) {}

  public usersExist(): boolean {
    return this.users.length > 0;
  }

  public deleteUser(userId: number): void {
    const r = this.adminApiRequester.deleteUserById(userId);
    r.then(result => {
      if (result.status === ResponseStatus.Ok) {
        this.users = this.users.filter(user => user.id !== userId);
        this.notifierService.send('Done');
      } else {
        this.notifierService.send(`${(result as ApiErrorResponse).error_detail}`);
        return Promise.reject((result as ApiErrorResponse).error_detail);
      }
    });
  }

  ngOnInit() {
    this.adminApiRequester.getUsersList()
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
