import styles from './EmptyWorkspace.module.css'

type EmptyWorkspaceProps = { onCreate: () => void }

const EmptyWorkspace = ({ onCreate }: EmptyWorkspaceProps) => (
  <div className={styles.emptyWorkspace}>
    <div className={styles.emptyArtwork}><span>✳</span><span>▤</span><span>◉</span></div>
    <div className={styles.emptyEyebrow}>A CLEAR SPACE TO THINK</div>
    <h1>Make room for<br />great work.</h1>
    <p>Create your first board and bring all the moving pieces into one calm, clear view.</p>
    <button className={styles.primaryButton} onClick={onCreate}><span>＋</span> Create your first board</button>
  </div>
)

export default EmptyWorkspace
