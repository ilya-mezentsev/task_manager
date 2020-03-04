import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {ApiUrlBuilder} from '../helpers/api-url-builder';
import {ApiDefaultResponse, ApiErrorResponse} from '../interfaces/api';
import {UsersListResponse, TasksListResponse} from '../interfaces/api-responses';

@Injectable({
  providedIn: 'root'
})
export class ApiRequesterService {
  private readonly groupsUsersEndpoint = '/group/lead/users';
  private  readonly tasksListWorkerEndpoint = '/group/worker/tasks';
  private  readonly taskLeadEndpoint = '/group/lead/task';
  private  readonly tasksListLeadEndpoint = 'group/lead/tasks';
  private readonly httpOptions = new HttpHeaders({'Content-Type': 'application/json'});

  constructor(
    private readonly http: HttpClient
  ) { }

  public async getUsersList(groupId: number): Promise<UsersListResponse | ApiErrorResponse> {
    const body = {group_id: groupId};
    return await this.http.post(
      ApiUrlBuilder.getApiUrlRequest(this.groupsUsersEndpoint),
      body
    ).toPromise() as UsersListResponse | ApiErrorResponse;
  }

  public async getTasksListByUser(userId: number): Promise<TasksListResponse | ApiErrorResponse> {
    const body = {
      user_id: userId
    };
    return await this.http.post(
      ApiUrlBuilder.getApiUrlRequest(this.tasksListWorkerEndpoint),
      body
    ).toPromise() as TasksListResponse | ApiErrorResponse;
  }

  public async getTasksListByGroup(groupId: number): Promise<TasksListResponse | ApiErrorResponse> {
    const body =  {
        group_id: groupId
    };
    return await this.http.post(
      ApiUrlBuilder.getApiUrlRequest(this.tasksListLeadEndpoint),
      body
    ).toPromise() as TasksListResponse | ApiErrorResponse;
  }

  public async assignTaskById(userId: number, taskId: number): Promise<ApiDefaultResponse | ApiErrorResponse> {
    const body = {
        user_id: userId,
        task: {id: taskId}
    };
    return await this.http.post(
      ApiUrlBuilder.getApiUrlRequest(this.taskLeadEndpoint),
      body
    ).toPromise() as ApiDefaultResponse | ApiErrorResponse;
  }

}
