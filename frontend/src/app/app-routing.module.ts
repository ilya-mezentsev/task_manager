import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {LoginComponent} from './login/login.component';
import {NotFoundComponent} from './not-found/not-found.component';
import {AdminComponent} from './admin/admin.component';
import {GroupComponent} from './group/group.component';
import {AuthGuard} from './auth.guard';
import {CreateGroupComponent} from './admin/create-group/create-group.component';
import {CreateUserComponent} from './admin/create-user/create-user.component';
import {CreateTasksComponent} from './admin/create-tasks/create-tasks.component';
import {GroupsListComponent} from './admin/groups-list/groups-list.component';
import {UsersListComponent} from './admin/users-list/users-list.component';
import {TasksListComponent} from './admin/tasks-list/tasks-list.component';

const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: 'login', component: LoginComponent },
  {
    path: 'admin',
    component: AdminComponent,
    canActivate: [AuthGuard],
    children: [
      { path: '', redirectTo: 'tasks-list', pathMatch: 'full' },
      { path: 'create-group', component: CreateGroupComponent },
      { path: 'create-user', component: CreateUserComponent },
      { path: 'create-tasks', component: CreateTasksComponent },
      { path: 'groups-list', component: GroupsListComponent },
      { path: 'users-list', component: UsersListComponent },
      { path: 'tasks-list', component: TasksListComponent },
      { path: '**', redirectTo: 'tasks-list' }
    ]
  },
  {
    path: 'group',
    component: GroupComponent,
    canActivate: [AuthGuard]
  },
  { path: '**', component: NotFoundComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
