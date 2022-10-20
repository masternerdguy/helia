import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { SignupComponent } from './auth/signup/signup.component';
import { LoginComponent } from './auth/login/login.component';
import { HomeComponent } from './home/home.component';
import { LoreComponent } from './lore/lore.component';
import { UnifiedPolicyComponent } from './unified-policy/unified-policy.component';
import { ForgotPasswordComponent } from './auth/forgot-password/forgot-password.component';

const routes: Routes = [
  { path: 'lore', component: LoreComponent },
  { path: 'policy', component: UnifiedPolicyComponent },
  { path: 'auth/signup', component: SignupComponent },
  { path: 'auth/signin', component: LoginComponent },
  { path: 'auth/forgot', component: ForgotPasswordComponent },
  { path: '', component: HomeComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule],
})
export class AppRoutingModule {}
