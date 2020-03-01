import { Component, OnInit } from '@angular/core';
import {AdminApiRequesterService} from '../../services/admin-api-requester.service';
import {ApiErrorResponse, ResponseStatus} from '../../interfaces/api';
import {NotifierService} from '../../services/notifier.service';

@Component({
  selector: 'app-create-tasks',
  templateUrl: './create-tasks.component.html',
  styleUrls: ['./create-tasks.component.scss']
})
export class CreateTasksComponent implements OnInit {

  constructor(
    private readonly adminApiRequester: AdminApiRequesterService,
    private readonly notifierService: NotifierService
  ) {}

  ngOnInit(): void {
  }

  public addTask(groupId: number, taskName: string, taskDescription: string): void {
    this.notifierService.send(`Trying to add task "${taskName}"`)
    const r = this.adminApiRequester.addNewTask(groupId, taskName, taskDescription);
    r.then(result => {
      if (result.status === ResponseStatus.Ok) {
        this.notifierService.send('Task added');
      } else {
        this.notifierService.send(`${(result as ApiErrorResponse).error_detail}`);
        return Promise.reject((result as ApiErrorResponse).error_detail);
      }
    });
  }

}
