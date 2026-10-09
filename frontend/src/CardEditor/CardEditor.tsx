import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import type { Card } from '../api'
import styles from './CardEditor.module.css'

const cardSchema = z.object({
  title: z.string().trim().min(1, 'Give this task a title.').max(150),
  description: z.string().max(1000, 'Keep the description under 1000 characters.'),
})

export type CardFormValues = z.infer<typeof cardSchema>

type CardEditorProps = {
  initial: Card | null
  busy: boolean
  onSubmit: (values: CardFormValues) => void
  onCancel: () => void
}

const CardEditor = ({ initial, busy, onSubmit, onCancel }: CardEditorProps) => {
  const { register, handleSubmit, formState: { errors } } = useForm<CardFormValues>({
    resolver: zodResolver(cardSchema),
    defaultValues: { title: initial?.title ?? '', description: initial?.description ?? '' },
  })

  return (
    <form className={styles.form} onSubmit={handleSubmit(onSubmit)}>
      <label>
        Task name
        <input autoFocus placeholder="What needs to get done?" {...register('title')} />
        {errors.title && <small>{errors.title.message}</small>}
      </label>
      <label>
        Description <span className={styles.optional}>Optional</span>
        <textarea rows={4} placeholder="Add a few details to help your team…" {...register('description')} />
        {errors.description && <small>{errors.description.message}</small>}
      </label>
      <div className={styles.modalActions}>
        <button type="button" className={styles.cancelButton} onClick={onCancel}>Cancel</button>
        <button className={styles.primaryButton} disabled={busy}>{busy ? 'Saving…' : initial ? 'Save changes' : 'Create task'}</button>
      </div>
    </form>
  )
}

export default CardEditor
