import {Component, OnInit} from '@angular/core';
import {ApiErrorResponse, ResponseStatus, UserRole} from '../../interfaces/api';
import {AdminApiRequesterService} from '../../services/admin-api-requester.service';
import {Task, TasksListResponse} from '../../interfaces/admin-api-responses';
import {NotifierService} from '../../services/notifier.service';

@Component({
  selector: 'app-all-tasks',
  templateUrl: './all-tasks.component.html',
  styleUrls: ['./all-tasks.component.scss']
})
export class AllTasksComponent implements OnInit {
  public userRole: UserRole = UserRole.Admin;
  public tasks: Task[] = [];

  constructor(
    private readonly adminApiRequester: AdminApiRequesterService,
    private readonly notifierService: NotifierService
  ) { }

  public deleteTask(taskId: number): void {
    console.log(`trying to delete task: id <${taskId}>`);
  }

  ngOnInit() {
    this.adminApiRequester.getTasksList()
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
