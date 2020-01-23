import {ApiDefaultResponse, ApiResponse} from './api';

export interface Group {
  name: string;
}

export type GroupsListResponse = ApiResponse<Group[]>;

export type deleteGroupResponse = ApiDefaultResponse;

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
