import { describe, it, expect, beforeEach, vi } from 'vitest'
import { render, waitFor, fireEvent } from '@testing-library/svelte'
import Chat from '../lib/Chat.svelte'
import { cockpit } from '../lib/stores/cockpit.svelte'
import { ListModelsForProvider, SupportedThinkLevels } from './mocks/wailsApp'

const switched: string[] = []

const baseProps = {
  task: { title: '', steps: [] } as never,
  messages: [] as never,
  awaitingReply: false,
  agentStatus: '',
  toolSteps: [] as never,
  streamingText: '',
  reasoningText: '',
  onSend: () => {},
  onSwitchProvider: async () => {},
  onSwitchThinkLevel: async () => {},
  onSwitchModel: async (m: string) => { switched.push(m) },
  onCancelPendingModel: async () => {},
  onSubmitAPIKey: async () => {},
  model: {
    provider: 'devin', modelName: 'swe-2-high',
    thinkLevel: 'max', approval: 'full-access', wireFormat: '',
  } as never,
}

async function openMenu() {
  const chip = document.querySelector('.model-chip') as HTMLButtonElement
  expect(chip).toBeTruthy()
  await fireEvent.click(chip)
}

beforeEach(() => {
  vi.clearAllMocks()
  vi.mocked(SupportedThinkLevels).mockResolvedValue(['medium', 'high', 'max'] as never)
  switched.length = 0
  Element.prototype.scrollIntoView = () => {}
  cockpit.chat = []
  cockpit.activeView = 'chat'
  vi.mocked(ListModelsForProvider).mockResolvedValue(['swe-2-high', 'claude-opus-5-5-medium'] as never)
})


describe('Devin composer', () => {
  it('shows a single effective effort and searches friendly names and IDs', async () => {
    render(Chat, baseProps as never)
    await openMenu()
    await waitFor(() => expect(document.querySelector('.model-chip')?.textContent).toContain('SWE-2 Max'))
    await waitFor(() => expect(document.querySelectorAll('.model-menu .mm-row')).toHaveLength(4))
    const rows = document.querySelectorAll('.model-menu .mm-row')
    await fireEvent.click(rows[2].querySelector('button')!)
    const search = document.querySelector('.model-search input') as HTMLInputElement
    expect(search).toBeTruthy()
    await fireEvent.input(search, { target: { value: 'Claude Opus' } })
    await waitFor(() => {
      const choices = document.querySelectorAll('.model-list .updrop-opt')
      expect(choices).toHaveLength(1)
      expect(choices[0].textContent).toContain('Claude Opus 5.5 Medium')
      expect(choices[0].querySelector('.d')?.textContent).toBe('claude-opus-5-5-medium')
    })
    await fireEvent.input(search, { target: { value: 'swe-2-high' } })
    expect(document.querySelectorAll('.model-list .updrop-opt')).toHaveLength(1)
    expect(document.querySelector('.model-list .updrop-opt .t')?.textContent).toBe('SWE-2 Max')
  })
})
