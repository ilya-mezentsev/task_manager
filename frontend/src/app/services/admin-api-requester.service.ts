import { Injectable } from '@angular/core';
import {GroupsListResponse, TasksListResponse} from '../interfaces/admin-api-responses';
import {HttpClient} from '@angular/common/http';
import {ApiUrlBuilder} from '../helpers/api-url-builder';
import {ApiErrorResponse} from '../interfaces/api';

@Injectable({
  providedIn: 'root'
})
export class AdminApiRequesterService {
  private readonly groupsListEndpoint = '/admin/groups';
  private readonly tasksListEndpoint = '/admin/tasks';

  constructor(
    private readonly http: HttpClient
  ) { }

  public async getGroupsList(): Promise<GroupsListResponse | ApiErrorResponse> {
    return await this.http.get(
      ApiUrlBuilder.getApiUrlRequest(this.groupsListEndpoint)
    ).toPromise() as GroupsListResponse | ApiErrorResponse;
  }

  public async getTasksList(): Promise<TasksListResponse | ApiErrorResponse> {
    return await this.http.get(
      ApiUrlBuilder.getApiUrlRequest(this.tasksListEndpoint)
    ).toPromise() as TasksListResponse | ApiErrorResponse;
  }
}
