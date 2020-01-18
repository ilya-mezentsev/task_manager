import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WorkerTasksListComponent } from './worker-tasks-list.component';

describe('WorkerTasksListComponent', () => {
  let component: WorkerTasksListComponent;
  let fixture: ComponentFixture<WorkerTasksListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WorkerTasksListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WorkerTasksListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
