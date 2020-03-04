import { Component, OnInit } from '@angular/core';
import {ApiErrorResponse, ResponseStatus, UserRole, UserSession} from '../../interfaces/api';
import {ApiRequesterService} from '../../services/api-requester.service';
import {Task, TasksListResponse} from '../../interfaces/api-responses';
import {NotifierService} from '../../services/notifier.service';
import {StorageService} from '../../services/storage.service';

@Component({
  selector: 'app-group-tasks-list',
  templateUrl: './group-tasks-list.component.html',
  styleUrls: ['./group-tasks-list.component.scss']
})
export class GroupTasksListComponent implements OnInit {
  public userRole: UserRole = UserRole.GroupLead;
  public user: UserSession = this.storage.getSession();
  public tasks: Task[] = [];

  constructor(
    private readonly apiRequester: ApiRequesterService,
    private readonly notifier: NotifierService,
    private readonly storage: StorageService
  ) { }

  public assignTask(taskId: number): void {
    const r = this.apiRequester.assignTaskById(0, taskId);
    this.notifier.send('Trying to assign task')
    r.then(result => {
      if (result.status === ResponseStatus.Ok) {
        this.notifier.send('Task assigned successfully');
      } else {
        this.notifier.send(`${(result as ApiErrorResponse).error_detail}`);
        return Promise.reject((result as ApiErrorResponse).error_detail);
      }
    });
  }

  ngOnInit() {
    this.apiRequester.getTasksListByGroup(this.user.group_id)
      .then(res => {
        if (res.status === ResponseStatus.Ok) {
          this.tasks = (res as TasksListResponse).data;
        } else {
          return Promise.reject((res as ApiErrorResponse).error_detail);
        }
      })
      .catch(err => console.log(err));
  }
}
