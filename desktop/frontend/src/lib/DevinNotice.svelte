<script lang="ts">
  import { t } from './i18n.svelte'
  import { BrowserOpenURL } from '../../wailsjs/runtime/runtime'
  import ProviderMark from './ProviderMark.svelte'
</script>

<aside class="devin-note">
  <strong><ProviderMark name="devin" size={14} /> Devin · {t('devin.agent')}</strong>
  <p>{t('devin.scope')}</p>
  <p>{t('devin.modes')}</p>
  <p>{t('devin.billing')}</p>
  <button class="ctrl tiny" onclick={() => BrowserOpenURL('https://docs.devin.ai/cli/index')}>{t('devin.setup')}</button>
</aside>

<style>
  .devin-note { padding: 14px; border: 1px solid var(--border, #333); border-radius: 12px; font-size: 12px; line-height: 1.6; text-align: left; }
  strong { display: inline-flex; align-items: center; gap: 6px; }
  p { margin: 6px 0 10px; opacity: .8; }
</style>
