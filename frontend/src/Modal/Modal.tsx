import { useEffect } from 'react'
import type { ReactNode } from 'react'
import styles from './Modal.module.css'

type ModalProps = {
  title: string
  onClose: () => void
  children: ReactNode
}

const Modal = ({ title, onClose, children }: ModalProps) => {
  useEffect(() => {
    const closeOnEscape = (event: KeyboardEvent) => {
      if (event.key === 'Escape') onClose()
    }

    window.addEventListener('keydown', closeOnEscape)
    return () => window.removeEventListener('keydown', closeOnEscape)
  }, [onClose])

  return (
    <div className={styles.modalBackdrop} onMouseDown={(event) => {
      if (event.target === event.currentTarget) onClose()
    }}>
      <section className={styles.modal} role="dialog" aria-modal="true" aria-labelledby="modal-title">
        <div className={styles.modalHeader}>
          <h2 id="modal-title">{title}</h2>
          <button type="button" className={styles.modalClose} onClick={onClose} aria-label="Close">×</button>
        </div>
        {children}
      </section>
    </div>
  )
}

export default Modal
