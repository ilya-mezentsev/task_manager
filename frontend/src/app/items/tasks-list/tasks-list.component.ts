import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {Task} from '../../interfaces/admin-api-responses';
import {UserRole} from '../../interfaces/api';

@Component({
  selector: 'app-tasks-list',
  templateUrl: './tasks-list.component.html',
  styleUrls: ['./tasks-list.component.scss']
})
export class TasksListComponent implements OnInit {
  @Input() public tasks: Task[];
  @Input() public userRole: UserRole;
  // admin
  @Output() public deleteTaskEmitter = new EventEmitter<number>();
  // group lead
  @Output() public assignTaskToEmitter = new EventEmitter<any>();
  // group worker
  @Output() public commentTaskEmitter = new EventEmitter<any>();
  @Output() public completeTaskEmitter = new EventEmitter<number>();

  constructor() { }

  public userId: number = 0;
  public taskComment: string = '';

  public deleteTask(taskId: number): void {
    this.deleteTaskEmitter.emit(taskId);
  }

  public assignTask(taskId: number): void {
    this.assignTaskToEmitter.emit([this.userId, taskId]);
  }

  public commentTask(taskId: number): void {
    this.commentTaskEmitter.emit([this.taskComment, taskId]);
  }

  public completeTask(taskId: number): void {
    this.completeTaskEmitter.emit(taskId);
  }

  public get admin(): UserRole {
    return UserRole.Admin;
  }

  public get groupLead(): UserRole {
    return UserRole.GroupLead;
  }

  public get groupWorker(): UserRole {
    return UserRole.GroupWorker;
  }

  ngOnInit() {
  }

}
