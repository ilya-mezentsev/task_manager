import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {UserRole} from '../../interfaces/api';

@Component({
  selector: 'app-task-actions',
  templateUrl: './task-actions.component.html',
  styleUrls: ['./task-actions.component.scss']
})
export class TaskActionsComponent implements OnInit {
  @Input() public userRole: UserRole;
  @Input() public taskId: number;
  @Output() public deleteTaskEmitter = new EventEmitter<number>();
  @Output() public assignTaskToWorkerEmitter = new EventEmitter<number>();
  @Output() public commentTaskEmitter = new EventEmitter<number>();
  @Output() public completeTaskEmitter = new EventEmitter<number>();

  constructor() { }

  public deleteTask(): void {
    this.deleteTaskEmitter.emit(this.taskId);
  }

  public assignTask(): void {
    this.assignTaskToWorkerEmitter.emit(this.taskId);
  }

  public commentTask(): void {
    this.commentTaskEmitter.emit(this.taskId);
  }

  public completeTask(): void {
    this.completeTaskEmitter.emit(this.taskId);
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
