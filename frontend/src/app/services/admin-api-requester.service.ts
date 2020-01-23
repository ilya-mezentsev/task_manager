import { Injectable } from '@angular/core';
import {GroupsListResponse, TasksListResponse, deleteGroupResponse} from '../interfaces/admin-api-responses';
import {HttpClient} from '@angular/common/http';
import {HttpHeaders} from '@angular/common/http';
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

  public async deleteGroupById(id: number): Promise<deleteGroupResponse | ApiErrorResponse> {
    const options = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
      body: {
        group_id: id
      }
    };

    return await this.http.delete(
      ApiUrlBuilder.getApiUrlRequest(this.groupsListEndpoint),
      options
    ).toPromise() as deleteGroupResponse | ApiErrorResponse;
  }

  public async getTasksList(): Promise<TasksListResponse | ApiErrorResponse> {
    return await this.http.get(
      ApiUrlBuilder.getApiUrlRequest(this.tasksListEndpoint)
    ).toPromise() as TasksListResponse | ApiErrorResponse;
  }
}
