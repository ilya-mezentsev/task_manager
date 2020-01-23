import {ApiDefaultResponse, ApiResponse} from './api';

export interface Group {
  id: number;
  name: string;
}

export type GroupsListResponse = ApiResponse<Group[]>;

export type DeleteGroupResponse = ApiDefaultResponse;

export interface Task {
  id: number;
  title: string;
  description: string;
  group_id: number;
  user_id: number;
  is_complete: boolean;
  comment: string;
}

export type TasksListResponse = ApiResponse<Task[]>;
