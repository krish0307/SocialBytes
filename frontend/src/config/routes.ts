import {
  Home as HomeIcon,
  EventNoteRounded as EventIcon,
  AccountBoxRounded as UserIcon,
  SettingsOutlined as SettingsIcon,
  ListAlt as ListIcon,
  SearchRounded as SearchIcon,
  LightbulbRounded as LightbulbRounded
} from '@mui/icons-material';
import { createEvent } from '@testing-library/react';

import { Home } from '../pages/Home';
import { Event } from '../pages/CreateEvent';
import  { SearchEvent }  from '../pages/SearchEvent';

import { Route } from '../types/Route';
import { EventPage } from '../pages/event';


const routes: Array<Route> = [
  {
    key: 'router-home',
    title: 'Home',
    description: 'Home',
    component: Home,
    path: '/',
    isEnabled: true,
    icon: HomeIcon,
    appendDivider: true,
  },
  {
    key: 'router-start-event',
    title: 'Start-Event',
    description: 'Create a new event',
    component: Event,
    path: '/start-event',
    isEnabled: true,
    icon: EventIcon,
  },
  {
    key: 'router-search',
    title: 'Search for Event',
    description: 'Event Search',
    path: '/search-event',
    component: SearchEvent,
    isEnabled: true,
    icon: SearchIcon,
    appendDivider: true,
  },
  {
    key: 'router-my-account',
    title: 'My Account',
    description: 'My Account',
    path: '/account',
    isEnabled: true,
    icon: UserIcon,
    subRoutes: [
      {
        key: 'router-settings',
        title: 'Settings',
        description: 'Account Settings',
        path: '/account/settings',
        isEnabled: true,
        icon: SettingsIcon,
      },
      {
        key: 'router-preferences',
        title: 'Preferences',
        description: 'Account Preferences',
        path: '/account/preferences',
        isEnabled: true,
        icon: ListIcon,
      },
    ],
  },
];

export default routes;
