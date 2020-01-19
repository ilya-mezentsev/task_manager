import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupUsersListComponent } from './group-users-list.component';

describe('GroupUsersListComponent', () => {
  let component: GroupUsersListComponent;
  let fixture: ComponentFixture<GroupUsersListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GroupUsersListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupUsersListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
