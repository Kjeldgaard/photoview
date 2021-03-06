import React from 'react'
import { createGlobalStyle } from 'styled-components'
import { Helmet } from 'react-helmet'
import Routes from './components/routes/Routes'
import Messages from './components/messages/Messages'
import { useTranslation } from 'react-i18next'
import { loadTranslations } from './localization'

const GlobalStyle = createGlobalStyle`
  * {
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  html {
    font-size: 0.85rem;
  }

  #root, body {
    height: 100%;
    margin: 0;
    font-size: inherit;
  }

  /* Make dimmer lighter */
  .ui.dimmer {
    background-color: rgba(0, 0, 0, 0.5);
  }
`

import 'semantic-ui-css/semantic.min.css'

const App = () => {
  const { t } = useTranslation()
  loadTranslations()

  return (
    <>
      <Helmet>
        <meta
          name="description"
          content={t(
            'meta.description',
            'Simple and User-friendly Photo Gallery for Personal Servers'
          )}
        />
      </Helmet>
      <GlobalStyle />
      <Routes />
      <Messages />
    </>
  )
}

export default App
