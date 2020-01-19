import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {LoginComponent} from './login/login.component';
import {NotFoundComponent} from './not-found/not-found.component';
import {AdminComponent} from './admin/admin.component';
import {GroupComponent} from './group/group.component';
import {AuthGuard} from './auth.guard';
import {CreateGroupComponent} from './admin/create-group/create-group.component';
import {CreateUserComponent} from './admin/create-user/create-user.component';
import {CreateTasksComponent} from './admin/create-tasks/create-tasks.component';
import {WorkerTasksListComponent} from './group/worker-tasks-list/worker-tasks-list.component';
import {GroupTasksListComponent} from './group/group-tasks-list/group-tasks-list.component';
import {GroupUsersListComponent} from './group/group-users-list/group-users-list.component';
import {UserRole} from './interfaces/api';
import {AllGroupsComponent} from './admin/all-groups/all-groups.component';
import {AllUsersComponent} from './admin/all-users/all-users.component';
import {AllTasksComponent} from './admin/all-tasks/all-tasks.component';

const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: 'login', component: LoginComponent },
  {
    path: 'admin',
    component: AdminComponent,
    canActivate: [AuthGuard],
    data: { roles: [UserRole.Admin] },
    children: [
      { path: '', redirectTo: 'tasks-list', pathMatch: 'full' },
      { path: 'create-group', component: CreateGroupComponent },
      { path: 'create-user', component: CreateUserComponent },
      { path: 'create-tasks', component: CreateTasksComponent },
      { path: 'groups-list', component: AllGroupsComponent },
      { path: 'users-list', component: AllUsersComponent },
      { path: 'tasks-list', component: AllTasksComponent },
      { path: '**', redirectTo: 'tasks-list' }
    ]
  },
  {
    path: 'group',
    component: GroupComponent,
    canActivate: [AuthGuard],
    data: { roles: [UserRole.GroupWorker, UserRole.GroupLead] },
    children: [
      { path: '', redirectTo: 'tasks-list', pathMatch: 'full' },
      { path: 'tasks-list', component: WorkerTasksListComponent },
      {
        path: 'group-users',
        component: GroupUsersListComponent,
        canActivate: [AuthGuard],
        data: { roles: [UserRole.GroupLead] }
      },
      {
        path: 'group-tasks',
        component: GroupTasksListComponent,
        canActivate: [AuthGuard],
        data: { roles: [UserRole.GroupLead] }
      },
      { path: '**', redirectTo: 'tasks-list' }
    ]
  },
  { path: '**', component: NotFoundComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
