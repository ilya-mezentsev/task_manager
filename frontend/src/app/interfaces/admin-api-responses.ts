import {ApiDefaultResponse, ApiResponse} from './api';

export interface Group {
  id: number;
  name: string;
}

export interface User {
  id: number;
  name: string;
  group_id: number;
  password: string;
  is_group_lead: boolean;
}

export interface Task {
  id: number;
  title: string;
  description: string;
  group_id: number;
  user_id: number;
  is_complete: boolean;
  comment: string;
}

export type GroupsListResponse = ApiResponse<Group[]>;

export type DeleteGroupResponse = ApiDefaultResponse;

export type addNewGroupResponse = ApiDefaultResponse;

export type UsersListResponse = ApiResponse<User[]>;

export type DeleteUserResponse = ApiDefaultResponse;

export type addNewUserResponse = ApiDefaultResponse;

export type TasksListResponse = ApiResponse<Task[]>;

export type DeleteTaskResponse = ApiDefaultResponse;

export type addNewTaskResponse = ApiDefaultResponse;
