export type Card = {
  id: string
  column_id: string
  title: string
  description: string
  order: number
}

export type Column = {
  id: string
  board_id: string
  title: string
  position: number
  cards?: Card[]
}

export type Board = {
  id: string
  title: string
  created_at?: string
  columns: Column[]
}

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const response = await fetch(path, {
    ...init,
    headers: { 'Content-Type': 'application/json', ...init?.headers },
  })
  if (!response.ok) {
    const message = await response.text()
    throw new Error(message || `Request failed (${response.status})`)
  }
  if (response.status === 204) return undefined as T
  return response.json() as Promise<T>
}

export const api = {
  boards: () => request<Board[]>('/boards'),
  board: (id: string) => request<Board>(`/boards/${id}`),
  createBoard: (title: string) => request<Board>('/boards', {
    method: 'POST', body: JSON.stringify({ id: crypto.randomUUID(), title, columns: [] }),
  }),
  createColumn: (boardId: string, title: string, position: number) => request<Column>('/columns', {
    method: 'POST', body: JSON.stringify({ id: crypto.randomUUID(), board_id: boardId, title, position }),
  }),
  createCard: (card: Omit<Card, 'id'>) => request<Card>('/cards', {
    method: 'POST', body: JSON.stringify({ ...card, id: crypto.randomUUID() }),
  }),
  updateCard: (id: string, card: Omit<Card, 'id'>) => request<Card>(`/cards/${id}`, {
    method: 'PUT', body: JSON.stringify({ ...card, id }),
  }),
  deleteCard: (id: string) => request<void>(`/cards/${id}`, { method: 'DELETE' }),
  moveCard: (id: string, columnId: string) => request<void>(`/cards/${id}/move`, {
    method: 'POST', body: JSON.stringify({ column_id: columnId }),
  }),
}
