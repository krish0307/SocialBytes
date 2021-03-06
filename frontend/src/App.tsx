import { useMemo, useState } from 'react';
import { CssBaseline, ThemeProvider } from '@mui/material';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';

import { Layout } from './components/Layout';
import  PageDefault  from './components/PageDefault';

import { AppContext, ThemeModeContext } from './contexts';
import { AppClient } from './clients';
import { routes } from './config';
import { Route as AppRoute } from './types';
import { getAppTheme } from './styles/theme';
import { DARK_MODE_THEME, LIGHT_MODE_THEME } from './utils/constants';
import { EventPage } from './pages/event';
import { SearchEvent } from './pages/SearchEvent';
import LoginComponent from './pages/LoginComponent';
import Register from './pages/Register';

function App() {
  const [mode, setMode] = useState<typeof LIGHT_MODE_THEME | typeof DARK_MODE_THEME>(DARK_MODE_THEME);
  const appClient = new AppClient();

  const themeMode = useMemo(
    () => ({
      toggleThemeMode: () => {
        setMode((prevMode) => (prevMode === LIGHT_MODE_THEME ? DARK_MODE_THEME : LIGHT_MODE_THEME));
      },
    }),
    []
  );

  const theme = useMemo(() => getAppTheme(mode), [mode]);

  const addRoute = (route: AppRoute) => (
    <Route key={route.key} path={route.path} component={route.component || PageDefault} exact />
  );

  return (
    <AppContext.Provider value={appClient}>
      <ThemeModeContext.Provider value={themeMode}>
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <Router>
            <Switch>
              <Layout>
                {routes.map((route: AppRoute) =>
                  route.subRoutes ? route.subRoutes.map((item: AppRoute) => addRoute(item)) : addRoute(route)
                )}
                 <Route key='eventpage' path='/eventpage/:id' component={EventPage}  />
                 <Route key='searchpage' path='/search-event/:query' component={SearchEvent}  />
                 <Route key='Login' path='/Login' component={LoginComponent}  />
                 <Route key='Register' path='/Register' component={Register}  />
              </Layout>
            </Switch>
          </Router>
        </ThemeProvider>
      </ThemeModeContext.Provider>
    </AppContext.Provider>
  );
}

export default App;