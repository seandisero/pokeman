package main

func commandPrintCache(cfg *config, args ...string) error {
	cfg.pokapiClient.PrintCache()
	return nil
}
