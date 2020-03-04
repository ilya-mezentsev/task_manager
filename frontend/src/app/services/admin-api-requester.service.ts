import { Injectable } from '@angular/core';
import {GroupsListResponse, TasksListResponse, DeleteGroupResponse, DeleteTaskResponse, UsersListResponse,
  DeleteUserResponse, addNewGroupResponse, addNewUserResponse, addNewTaskResponse} from '../interfaces/admin-api-responses';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {ApiUrlBuilder} from '../helpers/api-url-builder';
import {ApiErrorResponse} from '../interfaces/api';

@Injectable({
  providedIn: 'root'
})
export class AdminApiRequesterService {
  private readonly groupsListEndpoint = '/admin/groups';
  private readonly groupApiEndpoint = '/admin/group';
  private readonly tasksListEndpoint = '/admin/tasks';
  private readonly taskApiEndpoint = '/admin/task';
  private readonly usersListEndpoint = '/admin/users';
  private readonly userApiEndpoint = '/admin/user';
  private readonly httpOptions = new HttpHeaders({'Content-Type': 'application/json'});

  constructor(
    private readonly http: HttpClient
  ) { }

  public async getGroupsList(): Promise<GroupsListResponse | ApiErrorResponse> {
    return await this.http.get(
      ApiUrlBuilder.getApiUrlRequest(this.groupsListEndpoint)
    ).toPromise() as GroupsListResponse | ApiErrorResponse;
  }

  public async deleteGroupById(id: number): Promise<DeleteGroupResponse | ApiErrorResponse> {
    const options = {
      headers: this.httpOptions,
      body: {
        group_id: id
      }
    };

    return await this.http.delete(
      ApiUrlBuilder.getApiUrlRequest(this.groupApiEndpoint),
      options
    ).toPromise() as DeleteGroupResponse | ApiErrorResponse;
  }

  public async addNewGroup(groupName: string): Promise<addNewGroupResponse | ApiErrorResponse> {
    const body = {group_name: groupName};

    return await this.http.post(
      ApiUrlBuilder.getApiUrlRequest(this.groupApiEndpoint),
      body
    ).toPromise() as addNewGroupResponse | ApiErrorResponse;
  }

  public async deleteTaskById(id: number): Promise<DeleteTaskResponse | ApiErrorResponse> {
    const options = {
      headers: this.httpOptions,
      body: {
        task_id: id
      }
    };

    return await this.http.delete(
      ApiUrlBuilder.getApiUrlRequest(this.taskApiEndpoint),
      options
    ).toPromise() as DeleteGroupResponse | ApiErrorResponse;
  }

  public async getTasksList(): Promise<TasksListResponse | ApiErrorResponse> {
    return await this.http.get(
      ApiUrlBuilder.getApiUrlRequest(this.tasksListEndpoint)
    ).toPromise() as TasksListResponse | ApiErrorResponse;
  }

  public async getUsersList(): Promise<UsersListResponse | ApiErrorResponse> {
    return await this.http.get(
      ApiUrlBuilder.getApiUrlRequest(this.usersListEndpoint)
    ).toPromise() as UsersListResponse | ApiErrorResponse;
  }

  public async deleteUserById(id: number): Promise<DeleteUserResponse | ApiErrorResponse> {
    const options = {
      headers: this.httpOptions,
      body: {
        user_id: id
      }
    };

    return await this.http.delete(
      ApiUrlBuilder.getApiUrlRequest(this.userApiEndpoint),
      options
    ).toPromise() as DeleteUserResponse | ApiErrorResponse;
  }

  public async addNewUser(userName: string, groupId: number, isGroupLead: boolean): Promise<addNewUserResponse | ApiErrorResponse> {
    const body = {user: {
        name: userName,
        group_id: groupId,
        is_group_lead: isGroupLead
      }};
    return await this.http.post(
      ApiUrlBuilder.getApiUrlRequest(this.userApiEndpoint),
      body
    ).toPromise() as addNewUserResponse | ApiErrorResponse;
  }

  public async addNewTask(groupId: number, taskTitle: string, taskDescription: string): Promise<addNewTaskResponse | ApiErrorResponse> {
    const body = {
        group_id: groupId, tasks: [{title: taskTitle, description: taskDescription}]
    };
    return await this.http.post(
      ApiUrlBuilder.getApiUrlRequest(this.tasksListEndpoint),
      body
    ).toPromise() as addNewTaskResponse | ApiErrorResponse;
  }
}
