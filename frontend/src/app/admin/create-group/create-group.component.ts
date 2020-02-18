import { Component, OnInit } from '@angular/core';
import {AdminApiRequesterService} from '../../services/admin-api-requester.service';
import {Group} from '../../interfaces/admin-api-responses';
import {ApiErrorResponse} from '../../interfaces/api';
import {ResponseStatus} from '../../interfaces/api';
import {NotifierService} from '../../services/notifier.service';

@Component({
  selector: 'app-create-group',
  templateUrl: './create-group.component.html',
  styleUrls: ['./create-group.component.scss']
})
export class CreateGroupComponent implements OnInit {

  constructor(
    private readonly adminApiRequester: AdminApiRequesterService,
    private readonly notifierService: NotifierService
  ) {}

  public addGroup(groupName: string): void {
    this.notifierService.send(`Trying to add group "${groupName}"`)
    const r = this.adminApiRequester.addNewGroup(groupName);
    r.then(result => {
      if (result.status === ResponseStatus.Ok) {
        this.notifierService.send('Group added');
      } else {
        this.notifierService.send(`${(result as ApiErrorResponse).error_detail}`);
        return Promise.reject((result as ApiErrorResponse).error_detail);
      }
    });
  }

  ngOnInit() {
  }

}
