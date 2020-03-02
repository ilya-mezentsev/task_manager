import {Component, Input, OnInit} from '@angular/core';
import {ApiErrorResponse, ResponseStatus, UserSession} from '../../interfaces/api';
import {ApiRequesterService} from '../../services/api-requester.service';
import {Task, TasksListResponse} from '../../interfaces/api-responses';
import {NotifierService} from '../../services/notifier.service';
import {StorageService} from '../../services/storage.service';

@Component({
  selector: 'app-worker-tasks-list',
  templateUrl: './worker-tasks-list.component.html',
  styleUrls: ['./worker-tasks-list.component.scss']
})
export class WorkerTasksListComponent implements OnInit {
  public user: UserSession = this.storageService.getSession();
  public tasks: Task[] = [];

  constructor(
    private readonly apiRequesterService: ApiRequesterService,
    private readonly notifierService: NotifierService,
    private readonly storageService: StorageService
  ) { }

  ngOnInit() {
    this.apiRequesterService.getTasksList(this.user.id)
      .then(res => {
        if (res.status === ResponseStatus.Ok) {
          this.tasks = (res as TasksListResponse).data;
        } else {
          this.notifierService.send((res as ApiErrorResponse).error_detail);
          return Promise.reject((res as ApiErrorResponse).error_detail);
        }
      })
      .catch(err => console.log(err));
  }
}
