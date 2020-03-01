import { Component, OnInit } from '@angular/core';
import {AdminApiRequesterService} from '../../services/admin-api-requester.service';
import {ApiErrorResponse, ResponseStatus} from '../../interfaces/api';
import {NotifierService} from '../../services/notifier.service';

@Component({
  selector: 'app-create-user',
  templateUrl: './create-user.component.html',
  styleUrls: ['./create-user.component.scss']
})
export class CreateUserComponent implements OnInit {

  constructor(
    private readonly adminApiRequester: AdminApiRequesterService,
    private readonly notifierService: NotifierService
  ) {}

  isGroupLead = false;

  public addUser(userName: string, groupId: number, isGroupLead: boolean): void {
    this.notifierService.send(`Trying to add user "${userName}"`);
    const r = this.adminApiRequester.addNewUser(userName, groupId, isGroupLead);
    r.then(result => {
      if (result.status === ResponseStatus.Ok) {
        this.notifierService.send('User added');
      } else {
        this.notifierService.send(`${(result as ApiErrorResponse).error_detail}`);
        return Promise.reject((result as ApiErrorResponse).error_detail);
      }
    });
  }

  ngOnInit() {
  }
}
