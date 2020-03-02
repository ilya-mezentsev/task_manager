import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {ApiUrlBuilder} from '../helpers/api-url-builder';
import {ApiErrorResponse} from '../interfaces/api';
import {UsersListResponse, TasksListResponse} from '../interfaces/api-responses';

@Injectable({
  providedIn: 'root'
})
export class ApiRequesterService {
  private readonly groupsUsersEndpoint = '/group/lead/users';
  private  readonly tasksListEndpoint = '/group/worker/tasks';

  constructor(
    private readonly http: HttpClient
  ) { }

  public async getUsersList(groupId: number): Promise<UsersListResponse | ApiErrorResponse> {
    // const options = {
    //   headers: new HttpHeaders({
    //     'Content-Type': 'application/json'
    //   }),
    //   body: {
    //     group_id: groupId
    //   }
    // };
    const body = {group_id: groupId};
    return await this.http.post(
      ApiUrlBuilder.getApiUrlRequest(this.groupsUsersEndpoint),
      body
    ).toPromise() as UsersListResponse | ApiErrorResponse;
  }

  public async getTasksList(userId: number): Promise<TasksListResponse | ApiErrorResponse> {
    const options = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
      body: {
        user_id: userId
      }
    };
    return await this.http.post(
      ApiUrlBuilder.getApiUrlRequest(this.tasksListEndpoint),
      options
    ).toPromise() as TasksListResponse | ApiErrorResponse;
  }

}
