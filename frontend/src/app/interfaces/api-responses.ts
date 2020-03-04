import {ApiDefaultResponse, ApiResponse} from './api';

export interface User {
  id: number;
  name: string;
  group_id: number;
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

export type UsersListResponse = ApiResponse<User[]>;

export type TasksListResponse = ApiResponse<Task[]>;
