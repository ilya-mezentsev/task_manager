import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { NotFoundComponent } from './not-found/not-found.component';
import { AdminComponent } from './admin/admin.component';
import { GroupComponent } from './group/group.component';
import {FormsModule} from '@angular/forms';
import {HttpClientModule} from '@angular/common/http';
import { NavigationComponent as AdminNavigationComponent } from './admin/navigation/navigation.component';
import { NavigationComponent as GroupNavigationComponent } from './group/navigation/navigation.component';
import { CreateGroupComponent } from './admin/create-group/create-group.component';
import { CreateUserComponent } from './admin/create-user/create-user.component';
import { GroupsListComponent } from './admin/groups-list/groups-list.component';
import { UsersListComponent } from './admin/users-list/users-list.component';
import { TasksListComponent } from './admin/tasks-list/tasks-list.component';
import { CreateTasksComponent } from './admin/create-tasks/create-tasks.component';
import { AssignTasksComponent } from './group/assign-tasks/assign-tasks.component';
import { GroupTasksListComponent } from './group/group-tasks-list/group-tasks-list.component';
import { WorkerTasksListComponent } from './group/worker-tasks-list/worker-tasks-list.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    NotFoundComponent,
    AdminComponent,
    GroupComponent,
    AdminNavigationComponent,
    GroupNavigationComponent,
    CreateGroupComponent,
    CreateUserComponent,
    GroupsListComponent,
    UsersListComponent,
    TasksListComponent,
    CreateTasksComponent,
    AssignTasksComponent,
    GroupTasksListComponent,
    WorkerTasksListComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
