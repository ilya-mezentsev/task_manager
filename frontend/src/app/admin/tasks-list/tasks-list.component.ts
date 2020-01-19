import {Component, OnInit} from '@angular/core';
import {ApiErrorResponse, ResponseStatus, UserRole} from '../../interfaces/api';
import {AdminApiRequesterService} from '../../services/admin-api-requester.service';
import {Task, TasksListResponse} from '../../interfaces/admin-api-responses';

@Component({
  selector: 'app-tasks-list',
  templateUrl: './tasks-list.component.html',
  styleUrls: ['./tasks-list.component.scss']
})
export class TasksListComponent implements OnInit {
  public userRole: UserRole = UserRole.Admin;
  public tasks: Task[] = [];

  constructor(
    private readonly adminAdiRequester: AdminApiRequesterService
  ) { }

  public deleteTask(taskId: number): void {
    console.log(`trying to delete task: id <${taskId}>`);
  }

  ngOnInit() {
    this.adminAdiRequester.getTasksList()
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
