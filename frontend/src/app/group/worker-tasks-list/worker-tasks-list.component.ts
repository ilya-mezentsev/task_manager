import {Component, Input, OnInit} from '@angular/core';
import {ApiErrorResponse, ResponseStatus, UserRole, UserSession} from '../../interfaces/api';
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
  public user: UserSession = this.storage.getSession();
  public tasks: Task[] = [];
  public userRole: UserRole = UserRole.GroupWorker;

  constructor(
    private readonly apiRequester: ApiRequesterService,
    private readonly notifier: NotifierService,
    private readonly storage: StorageService
  ) { }

  ngOnInit() {
    this.apiRequester.getTasksListByUser(this.user.id)
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
