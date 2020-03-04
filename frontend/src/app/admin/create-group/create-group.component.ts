import { Component, OnInit } from '@angular/core';
import {AdminApiRequesterService} from '../../services/admin-api-requester.service';
import {ApiErrorResponse, ResponseStatus} from '../../interfaces/api';
import {NotifierService} from '../../services/notifier.service';

@Component({
  selector: 'app-create-group',
  templateUrl: './create-group.component.html',
  styleUrls: ['./create-group.component.scss']
})
export class CreateGroupComponent implements OnInit {

  constructor(
    private readonly adminApi: AdminApiRequesterService,
    private readonly notifier: NotifierService
  ) {}

  public addGroup(groupName: string): void {
    this.notifier.send(`Trying to add group "${groupName}"`)
    const r = this.adminApi.addNewGroup(groupName);
    r.then(result => {
      if (result.status === ResponseStatus.Ok) {
        this.notifier.send('Group added');
      } else {
        this.notifier.send(`${(result as ApiErrorResponse).error_detail}`);
        return Promise.reject((result as ApiErrorResponse).error_detail);
      }
    });
  }

  ngOnInit() {
  }

}
