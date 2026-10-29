import { describe, expect, it } from 'vitest'
import { devinChoices, devinLabel } from './devinModels'
describe('Devin model picker', () => {
  it('labels canonical and unknown IDs without losing billing identity', () => {
    expect(devinLabel('swe-2-lightning-max')).toBe('SWE-2 Max (Lightning)')
    expect(devinLabel('swe-2-priority-high')).toBe('SWE-2 High (Priority)')
    expect(devinLabel('devin-swe-1-7-paid')).toBe('SWE-1.7 Priority')
    expect(devinLabel('claude-opus-4-5')).toBe('Claude Opus 4.5')
    expect(devinLabel('gpt-6-1-sol')).toBe('GPT 6.1 Sol')
  })
  it('shows one effective tier when ACP overrides the base model', () => {
    expect(devinLabel('swe-2-high', 'max')).toBe('SWE-2 Max')
    expect(devinChoices(['swe-2-high'], 'swe-2-high', 'max')[0].label).toBe('SWE-2 Max')
    expect(devinLabel('claude-opus-5-5-medium', 'high')).toBe('Claude Opus 5.5 High')
  })
  it('groups only returned family members and retains the chosen tier', () => {
    const choices = devinChoices(['swe-2-medium', 'swe-2-max', 'claude-opus-4-5'], 'swe-2-max')
    expect(choices).toHaveLength(2)
    expect(choices[0]).toEqual({ value: 'swe-2-max', label: 'SWE-2 Max', desc: 'swe-2-medium · swe-2-max' })
    expect(choices[1].label).toBe('Claude Opus 4.5')
  })
})
