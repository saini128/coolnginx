# 📝 TODO List for Nginx Config Manager  

## ✅ **Core Features**  
- [x] Load Nginx configuration using `crossplane`.  
- [x] Store Nginx configuration in **BoltDB** with a timestamp.  
- [x] Maintain only **last 5 stored instances** (FIFO mechanism).  
- [x] Prevent storing duplicate configurations (return error if unchanged).  
- [ ] Fetch all the virtual hosts from the configurations and store them in a bucket.


## 🚀 **TUI Interface**  
- [ ] Init **CLI/TUI interface** Dashboard (Logo, List Configs, Add Config, Ask AI, Help Menu).  
- [ ] Setup **List COnfig** and show all the server blocks.  
- [ ] Implement **search/filtering** for stored configurations.  
- [ ] Add **Modify** functionality for configurations.  
- [ ] Add **Create** functionality for configurations.
- [ ] Create **Help Menu** and list all the functionalities and how to use them.
- [ ] Make **Ask AI** and implement chatbot features in it with configuration file as context.



## 🔧 **Error Handling & Optimization**  
- [ ] Improve error messages for better debugging.  
- [ ] Optimize key management in **BoltDB** (use timestamps instead of index).  
- [ ] Implement **concurrent access handling** for BoltDB.  

## 📦 **Packaging & Deployment**  
- [ ] Create a **cross-platform installer** (MSI for Windows, DEB for Linux, PKG for macOS).  
- [ ] Ensure **zero setup**—users should be able to run immediately after installation.  
- [ ] Provide **automatic updates** for future improvements.  

## 🛠 **Testing & Validation**  
- [ ] Write **unit tests** for configuration parsing & storage.  
- [ ] Test compatibility with **various Nginx versions**.  
- [ ] Validate performance with **large configurations**.  
