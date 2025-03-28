// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { UserService } from './user.service';
import { MessageHandlerService } from '../../../shared/services/message-handler.service';
import { AppConfigService } from '../../../services/app-config.service';
import { SessionService } from '../../../shared/services/session.service';
import { UserComponent } from './user.component';
import { OperationService } from '../../../shared/components/operation/operation.service';
import { ConfirmationDialogService } from '../../global-confirmation-dialog/confirmation-dialog.service';
import { SharedTestingModule } from '../../../shared/shared.module';

describe('UserComponent', () => {
    let component: UserComponent;
    let fixture: ComponentFixture<UserComponent>;
    let fakeSessionService = null;
    let fakeAppConfigService = {
        getConfig: function () {
            return {
                auth_mode: 'ldap_auth',
            };
        },
    };
    let fakeUserService = null;
    let fakeMessageHandlerService = {
        handleError: function () {},
    };

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [UserComponent],
            imports: [SharedTestingModule],
            schemas: [CUSTOM_ELEMENTS_SCHEMA],
            providers: [
                ConfirmationDialogService,
                OperationService,
                {
                    provide: MessageHandlerService,
                    useValue: fakeMessageHandlerService,
                },
                { provide: UserService, useValue: fakeUserService },
                { provide: SessionService, useValue: fakeSessionService },
                { provide: AppConfigService, useValue: fakeAppConfigService },
            ],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(UserComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
