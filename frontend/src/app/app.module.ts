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
import { AllGroupsComponent } from './admin/all-groups/all-groups.component';
import { AllUsersComponent } from './admin/all-users/all-users.component';
import { AllTasksComponent } from './admin/all-tasks/all-tasks.component';
import { CreateTasksComponent } from './admin/create-tasks/create-tasks.component';
import { GroupTasksListComponent } from './group/group-tasks-list/group-tasks-list.component';
import { WorkerTasksListComponent } from './group/worker-tasks-list/worker-tasks-list.component';
import { GroupUsersListComponent } from './group/group-users-list/group-users-list.component';
import { LogoutComponent } from './share/logout/logout.component';
import { TaskActionsComponent } from './share/task-actions/task-actions.component';

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
    AllGroupsComponent,
    AllUsersComponent,
    AllTasksComponent,
    CreateTasksComponent,
    GroupTasksListComponent,
    WorkerTasksListComponent,
    GroupUsersListComponent,
    LogoutComponent,
    TaskActionsComponent
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
