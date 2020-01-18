import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupTasksListComponent } from './group-tasks-list.component';

describe('GroupTasksListComponent', () => {
  let component: GroupTasksListComponent;
  let fixture: ComponentFixture<GroupTasksListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GroupTasksListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupTasksListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
